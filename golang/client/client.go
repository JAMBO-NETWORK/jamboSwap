package client

import (
	"context"
	"github.com/astaxie/beego/logs"
	"github.com/ethereum/go-ethereum/ethclient"
	"time"
)

func CreateContractInstance(contractAddress string) {

}

func CreateClient(clientUrl string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(clientUrl)
	if err != nil {
		logs.Error("ethclient.Dial error")
		return nil, err
	}
	defer client.Close()
	return client, nil
}

func timeoutContext() context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	go func() {
		time.Sleep(time.Second * 60)
		cancel()
	}()
	return ctx
}
