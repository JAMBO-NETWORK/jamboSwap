package mining

import (
	"github.com/astaxie/beego/logs"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// 获取每区块产币量
func GetTokenPerBlock(miningInstance *Mining) (*big.Int, error) {
	zero := big.NewInt(0)
	yamPerBlock, err := miningInstance.YamPerBlock(&bind.CallOpts{})
	if err != nil {
		logs.Error("miningInstance.YamPerBlock error: ", err)
		return zero, err
	}
	return yamPerBlock, nil
}

// 获取每区块产币量
func GetTotalAllocPoint(miningInstance *Mining) (*big.Int, error) {
	zero := big.NewInt(0)
	totalAllocPoint, err := miningInstance.TotalAllocPoint(&bind.CallOpts{})
	if err != nil {
		logs.Error("miningInstance.TotalAllocPoint error: ", err)
		return zero, err
	}
	return totalAllocPoint, nil
}

// 获取流动池信息
func GetPoolInfo(miningInstance *Mining, poolIndex int) (struct {
	LpToken         common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccYamPerShare  *big.Int
	TotalAmount     *big.Int
}, error) {
	return miningInstance.PoolInfo(&bind.CallOpts{}, big.NewInt(int64(poolIndex)))
}
