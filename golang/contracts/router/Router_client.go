package router

import (
	"github.com/astaxie/beego/logs"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// 调用pancakeRouter02合约中的getAmountsOut方法
func GetAmountsOutGo(routerInstance *Router, amountIn *big.Int, path []string) (*big.Int, error) {
	zero := big.NewInt(0)
	pathLen := len(path)
	path2 := make([]common.Address, pathLen)
	for i := 0; i < pathLen; i++ {
		path2[i] = common.HexToAddress(path[i])
	}
	amountsOut, err := routerInstance.GetAmountsOut(&bind.CallOpts{}, amountIn, path2)
	if err != nil {
		logs.Error("routerInstance.GetAmountsOut error: ", err)
		return zero, err
	}
	return amountsOut[len(amountsOut)-1], err
}
