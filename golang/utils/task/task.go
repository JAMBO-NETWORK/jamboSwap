package task

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/toolbox"
)

// 初始化定时任务
func Init() {
	// 链上链下容错，如果有效广告没有对应链上的抵押资产，则将此广告改为无效
	checkOnChainStatus := toolbox.NewTask("checkOnChainStatus", "0 0 0/1 * * *", CheckOnChainStatusPanic)
	toolbox.AddTask("checkOnChainStatus", checkOnChainStatus)

	// 从合约中获取流动池信息
	listenLiquidityInfo := toolbox.NewTask("listenLiquidityInfo", "0/30 * * * * *", ListenLiquidityInfo)
	toolbox.AddTask("listenLiquidityInfo", listenLiquidityInfo)

	toolbox.StartTask()
	logs.Info("Timed Task Init Successfully!")
}
