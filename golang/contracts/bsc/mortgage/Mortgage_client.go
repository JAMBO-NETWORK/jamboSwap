package mortgage

import (
	"github.com/astaxie/beego/logs"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"jamSwap/client"
	"math/big"
)

// 获取Mortgage实例
func GetMortgageInstance(clientUrl, mortgageAddress string) (*Mortgage, error) {
	client, err := client.CreateClient(clientUrl)
	if err != nil {
		logs.Error("client.CreateClient("+clientUrl+") is err: ", err)
		return nil, err
	}
	instance, err := NewMortgage(common.HexToAddress(mortgageAddress), client)
	if err != nil {
		logs.Error("NewMortgage error: ", err)
		return nil, err
	}
	return instance, nil
}

// 获取用户质押的代币
func GetMortgageAmount(mortgageInstance *Mortgage, userAddr string) (*big.Int, error) {
	zero := big.NewInt(0)
	userStruct, err := mortgageInstance.UserMap(&bind.CallOpts{}, common.HexToAddress(userAddr))
	if err != nil {
		logs.Error("mortgageInstance.UserMap error: ", err)
		return zero, err
	}
	return userStruct.Amount, nil
}
