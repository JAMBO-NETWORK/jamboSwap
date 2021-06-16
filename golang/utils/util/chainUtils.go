package util

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"jamSwap/contracts/bsc/mining"
	"jamSwap/contracts/pair"
	"jamSwap/contracts/router"
	decimalUtils "jamSwap/utils/decimals"
	"math/big"
	"strings"
)

var base18 int64 = 1000000000000000000

// 获取每个LP代币价值多少USDT
func GetLpPrice(chainType string, client *ethclient.Client, routerInstance *router.Router, lpToken string, isLp int) (*big.Float, error) {
	if isLp == 0 { // 单币种
		if IsUsdt(lpToken, chainType) {
			return big.NewFloat(1), nil
		}
		lpTokenPriceBig, err := router.GetAmountsOutGo(routerInstance, big.NewInt(base18), []string{lpToken, GetChainUsdtAddr(chainType)})
		if err != nil {
			return big.NewFloat(0), err
		}
		return decimalUtils.RemoveDecimals(lpTokenPriceBig, 18), nil
	} else { // 双币种
		lpTokenAddress := common.HexToAddress(lpToken)
		lpInstance, err := pair.NewPair(lpTokenAddress, client)
		if err != nil {
			return nil, err
		}
		totalSupply, err := lpInstance.TotalSupply(&bind.CallOpts{})
		if err != nil {
			return nil, err
		}
		logs.Info("lpToken=", lpToken, "中的totalSupply=", totalSupply)
		/*token0, err := lpInstance.Token0(&bind.CallOpts{})
		if err != nil {
			return nil, err
		}
		token0Instance, err := pair.NewPair(token0, client)
		if err != nil {
			return nil, err
		}
		token0Balance, err := token0Instance.BalanceOf(&bind.CallOpts{}, lpTokenAddress)
		if err != nil {
			return nil, err
		}*/
		//fmt.Println("token0Balance: ", token0Balance)

		token1, err := lpInstance.Token1(&bind.CallOpts{})
		if err != nil {
			return nil, err
		}
		token1Instance, err := pair.NewPair(token1, client)
		if err != nil {
			return nil, err
		}
		token1Balance, err := token1Instance.BalanceOf(&bind.CallOpts{}, lpTokenAddress)
		if err != nil {
			return nil, err
		}
		logs.Info("lpToken=", lpToken, "中的token1Balance=", token1Balance)

		/*token0Amount, err := getTokenAmount(routerInstance, token0.String(), chainType, token0Balance, totalSupply)
		if err != nil {
			return big.NewFloat(0), err
		}
		logs.Info("lpToken=", lpToken, "中的token0Amount=", token0Amount)*/

		token1Amount, err := getTokenAmount(routerInstance, token1.String(), chainType, token1Balance, totalSupply)
		if err != nil {
			return big.NewFloat(0), err
		}
		logs.Info("lpToken=", lpToken, "中的token1Amount=", token1Amount)
		return decimalUtils.Mul(token1Amount, big.NewFloat(2)), nil
	}
}

// 每日产出总价值USDT
func GetTotalTokenPerDayUSDT(miningInstance *mining.Mining, jamPrice *big.Float, chainType string) (*big.Float, error) {
	tokenPerBlockBig, err := mining.GetTokenPerBlock(miningInstance)
	if err != nil {
		return nil, err
	}
	tokenPerBlock := decimalUtils.RemoveDecimals(tokenPerBlockBig, 18)
	blockNumberPerDay, _ := new(big.Float).SetString(beego.AppConfig.String(chainType + "::blockNumberPerDay"))
	// 每日产出总价值USDT
	totalTokenPerDayUSDT := decimalUtils.Mul(decimalUtils.Mul(tokenPerBlock, blockNumberPerDay), jamPrice)
	return totalTokenPerDayUSDT, nil
}

// 获取tokenBalance1数量的token价值多少USDT
func getTokenAmount(routerInstance *router.Router, token, chainType string, tokenBalance1, totalSupply1 *big.Int) (*big.Float, error) {
	tokenAmount := big.NewFloat(0)
	tokenBalance := decimalUtils.RemoveDecimals(tokenBalance1, 18)
	totalSupply := decimalUtils.RemoveDecimals(totalSupply1, 18)
	tokenAmount1 := decimalUtils.Div(tokenBalance, totalSupply)
	if IsUsdt(token, chainType) { // 如果是USDT直接返回数量
		tokenAmount = tokenAmount1
	} else {
		// 在dex中获取token价格，默认BSC中用的BUSD，火币中用的HUSD，如果要用USDT要做调整
		tokenPriceBig, err := router.GetAmountsOutGo(routerInstance, big.NewInt(base18),
			[]string{token, GetChainUsdtAddr(chainType)})
		if err != nil {
			return big.NewFloat(0), err
		}
		tokenAmount = big.NewFloat(0).Mul(tokenAmount1, decimalUtils.RemoveDecimals(tokenPriceBig, 18))
	}
	return tokenAmount, nil
}

// 判断代币是否是USDT或者BUSDT
func IsUsdt(token string, chainType string) bool {
	if strings.EqualFold(token, GetChainUsdtAddr(chainType)) || strings.EqualFold(token, GetUsdtAddr(chainType)) {
		return true
	} else {
		return false
	}
}
func IsJam(token string, chainType string) bool {
	if strings.EqualFold(token, GetJamAddr(chainType)) {
		return true
	} else {
		return false
	}
}

func GetChainUsdtAddr(chainType string) string {
	key := ""
	if strings.EqualFold(chainType, "BSC") {
		key = chainType + "::BUSD"
	} else if strings.EqualFold(chainType, "HECO") {
		key = chainType + "::HUSD"
	}
	return beego.AppConfig.String(key)
}

func GetUsdtAddr(chainType string) string {
	return beego.AppConfig.String(chainType + "::USDT")
}
func GetJamAddr(chainType string) string {
	return beego.AppConfig.String(chainType + "::JAM")
}
