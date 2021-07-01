package task

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/ethereum/go-ethereum/common"
	"jamSwap/client"
	"jamSwap/contracts/bsc/mining"
	"jamSwap/contracts/bsc/mortgage"
	"jamSwap/contracts/router"
	"jamSwap/mvc/dao"
	decimalUtils "jamSwap/utils/decimals"
	"jamSwap/utils/util"
	"math/big"
	"runtime/debug"
	"strconv"
	"strings"
)

var base18 int64 = 1000000000000000000

func CheckOnChainStatusPanic() error {
	fmt.Println("11111111111111111111111")
	defer func() {
		if p := recover(); p != nil {
			logs.Error("panic recover! p:", p)
			debug.PrintStack()
		}
	}()
	return CheckOnChainStatus()
}

// 链上链下容错系统，如果有效广告没有对应链上的抵押资产，则将此广告改为无效
func CheckOnChainStatus() error {
	var dao dao.UserDao
	list := dao.QueryValidList()
	if len(list) > 0 {
		for i := 0; i < len(list); i++ {
			// 默认的四张图片
			if list[i].Id == 1 || list[i].Id == 2 || list[i].Id == 3 || list[i].Id == 4 {
				continue
			}
			userAddr := list[i].UserAddr
			chainType := list[i].ChainType
			// 查询用户的链上信息，余额是否大于mortgageAmount
			mortgageAmount, _ := beego.AppConfig.Int("home::mortgageAmount")
			clientUrl := beego.AppConfig.String(chainType + "::clientUrl")
			mortgageAddress := beego.AppConfig.String(chainType + "::mortgageAddress")
			mortgageInstance, err := mortgage.GetMortgageInstance(clientUrl, mortgageAddress)
			if err != nil {
				logs.Error("GetMortgageInstance error: ", err)
				return err
			}
			chainMortgageAmountBig, err := mortgage.GetMortgageAmount(mortgageInstance, userAddr)
			if err != nil {
				continue
			}
			chainMortgageAmountBig.Div(chainMortgageAmountBig, big.NewInt(base18))
			chainMortgageAmount := int(chainMortgageAmountBig.Int64())
			if chainMortgageAmount < mortgageAmount {
				// 无效广告
				logs.Warn("id=", list[i].Id, "，userAddr=", userAddr+"，chainType=", chainType,
					"的广告在链上的抵押资产为", chainMortgageAmount, "，是无效广告，删除。")
				dao.UpdateAdvertisement(userAddr, chainType, 2)
			}
		}
	}
	return nil
}

// 从合约中获取流动池信息
func ListenLiquidityInfo() error {
	return UpdateLiquidityInfoPanic("BSC")
}

func UpdateLiquidityInfoPanic(chainType string) error {
	defer func() {
		if p := recover(); p != nil {
			logs.Error("panic recover! p:", p)
			debug.PrintStack()
		}
	}()
	return UpdateLiquidityInfo(chainType)
}

// 从合约中获取流动池信息
func UpdateLiquidityInfo(chainType string) error {
	// 获取jam价格
	clientUrl := beego.AppConfig.String(chainType + "::clientUrl")
	client, err := client.CreateClient(clientUrl)
	if err != nil {
		logs.Error("client.CreateClient("+clientUrl+") is err: ", err)
		return err
	}
	// TODO
	// 上线记得修改
	//var routerInstance *router.Router
	routerAddress := beego.AppConfig.String(chainType + "::routerAddress")
	routerInstance, err := router.NewRouter(common.HexToAddress(routerAddress), client)
	if err != nil {
		logs.Error("NewMortgage error: ", err)
		return err
	}
	JAMAddr := beego.AppConfig.String(chainType + "::JAM")
	jamPrice, err := util.GetLpPrice(chainType, client, routerInstance, JAMAddr, 0)
	if err != nil {
		logs.Error("get jamPrice error: ", err)
		return err
	}
	//jamPrice := big.NewFloat(0.2)
	logs.Info("jamPrice: ", jamPrice)

	miningAddress := beego.AppConfig.String(chainType + "::miningAddress")
	miningInstance, err := mining.NewMining(common.HexToAddress(miningAddress), client)
	if err != nil {
		logs.Error("NewMining error: ", err)
		return err
	}
	// 获取每区块产币量总价值多少USDT
	totalTokenPerDayUSDT, err := util.GetTotalTokenPerDayUSDT(miningInstance, jamPrice, chainType)
	if err != nil {
		logs.Error("mining.GetTokenPerBlock error: ", err)
		return err
	}
	// 总分配点数
	totalAllocPoint, err := mining.GetTotalAllocPoint(miningInstance)
	if err != nil {
		logs.Error("mining.GetTotalAllocPoint error: ", err)
		return err
	}

	tvlNum := big.NewFloat(0)    // 所有流动池的总TVL数量（币数量）
	tvlAmount := big.NewFloat(0) // 所有流动池的总TVL USDT（价值USDT的数量）
	// 获取所有流动池
	var dao dao.LiquidityDao
	list := dao.LiquidityList(chainType)
	for i := 0; i < len(list); i++ {
		liquidityInfo := list[i]
		lpId := liquidityInfo.LpId
		isLp := liquidityInfo.IsLp
		totalLiquidityNew := liquidityInfo.TotalLiquidity
		lpPriceNew := liquidityInfo.LpPrice
		totalAmountNew := liquidityInfo.TotalAmount
		apyNew := liquidityInfo.Apy

		poolInfo, err := mining.GetPoolInfo(miningInstance, lpId)
		if err != nil {
			logs.Error("mining.GetPoolInfo error: ", err)
			continue
		}
		lpToken := poolInfo.LpToken.String()
		lpAddress := liquidityInfo.LpAddress
		if !strings.EqualFold(lpAddress, lpToken) {
			logs.Error("注意危险：lpId=", lpId, "的流动性池，数据库中的地址（", lpAddress, "）与合约中的地址（", lpToken, "）不一致！！！")
			continue
		}

		// 获取每个LP代币价值多少USDT
		lpPriceNew1 := big.NewFloat(0)
		if util.IsJam(lpToken, chainType) {
			lpPriceNew1 = jamPrice
		} else {
			lpPriceNew1, err = util.GetLpPrice(chainType, client, routerInstance, lpAddress, isLp)
			if err != nil {
				logs.Error("lpId=", lpId, "的util.GetLpPrice error: ", err)
				continue
			}
		}
		lpPriceNew = lpPriceNew1.String()
		logs.Info("lpId=", lpId, "，的价格(USDT): ", lpPriceNew)
		// 有质押量
		if poolInfo.TotalAmount.Cmp(big.NewInt(0)) > 0 {
			totalLiquidityNew1 := decimalUtils.RemoveDecimals(poolInfo.TotalAmount, 18)
			// totalLiquidityNew1 := big.NewFloat(1)
			totalLiquidityNew = totalLiquidityNew1.String()
			totalAmountNew1 := big.NewFloat(0).Mul(totalLiquidityNew1, lpPriceNew1)
			totalAmountNew = totalAmountNew1.String()
			// 获取该流动池每日产出的jam总价值多少USDT
			tokenPerDayUSDTLPAmount := getTokenPerDayUSDTLPAmount(poolInfo.AllocPoint, totalAllocPoint, totalTokenPerDayUSDT)
			// 该流动池总质押量
			totalAmountNewBigFloat, _ := big.NewFloat(0).SetString(totalAmountNew)
			// 年华收益率
			apyNewF, _ := decimalUtils.Div(decimalUtils.Mul(decimalUtils.Mul(tokenPerDayUSDTLPAmount, big.NewFloat(365)), big.NewFloat(100)),
				totalAmountNewBigFloat).Float64()
			apyNew = strconv.FormatFloat(decimalUtils.Decimal(apyNewF), 'f', -1, 64)
			// 总tvl
			tvlNum = decimalUtils.Add(tvlNum, totalLiquidityNew1)
			tvlAmount = decimalUtils.Add(tvlAmount, totalAmountNew1)
		} else {
			totalLiquidityNew = "0"
			totalAmountNew = "0"
			apyNew = "0"
		}

		logs.Info("lpId=", lpId, "的更新信息：totalLiquidityNew=", totalLiquidityNew, "，lpPriceNew=", lpPriceNew,
			"，totalAmountNew=", totalAmountNew, "，apyNew=", apyNew)
		// 更新流动池信息
		updateLiquidity(liquidityInfo.Id, totalLiquidityNew, lpPriceNew, totalAmountNew, apyNew)
		// 更新总流动性和jamPrice
		updateAlls(tvlNum, tvlAmount, jamPrice, chainType)
	}
	return nil
}

func updateAlls(tvlNum, tvlAmount, jamPrice *big.Float, chainType string) {
	var dao dao.ConfigDao
	dao.UpdateConfig("tvlNum", tvlNum.String(), chainType)
	dao.UpdateConfig("tvlAmount", tvlAmount.String(), chainType)
	dao.UpdateConfig("jamPrice", jamPrice.String(), chainType)
}

func updateLiquidity(id int, totalLiquidityNew, lpPriceNew, totalAmountNew, apyNew string) {
	var dao dao.LiquidityDao
	_, err := dao.Update(id, totalLiquidityNew, lpPriceNew, totalAmountNew, apyNew)
	if err != nil {
		logs.Error("updateLiquidity error: ", err)
	}
}

// 获取该流动池每日产出的jam总价值多少USDT
func getTokenPerDayUSDTLPAmount(allocPoint, totalAllocPoint *big.Int, totalTokenPerDayUSDT *big.Float) *big.Float {
	allocPoint1 := decimalUtils.RemoveDecimals(allocPoint, 0)
	totalAllocPoint1 := decimalUtils.RemoveDecimals(totalAllocPoint, 0)
	tokenPerDayUSDTLPAmount := decimalUtils.Div(big.NewFloat(0).Mul(totalTokenPerDayUSDT, allocPoint1), totalAllocPoint1)
	return tokenPerDayUSDTLPAmount
}
