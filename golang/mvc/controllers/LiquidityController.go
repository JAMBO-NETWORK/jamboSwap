package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"jamSwap/mvc/bo"
	"jamSwap/mvc/service"
	"jamSwap/utils/task"
	"jamSwap/utils/util"
)

type LiquidityController struct {
	beego.Controller
}

// 获取所有流动池的总流动性和jamPrice
func (liq *LiquidityController) GetTotalTvlAndJamPrice() {
	liq.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", liq.Ctx.Request.Header.Get("Origin"))
	var response bo.ResponseBo
	response.Code = "00001"
	response.Message = "fail"

	chainType := liq.GetString("chainType")
	var service service.LiquidityService
	response, err := service.GetTotalTvlAndJamPrice(chainType)
	if err != nil {
		logs.Error("service.GetTotalTvlAndJamPrice() is err: ", err)
	}
	liq.Data["json"] = response
	liq.ServeJSON()
}

// 获取流动池信息
// @param lpId 流动池id
// @param chainType 链类型：ETH,HECO,BSC,DOT
// @return 返回流动池的信息
func (liq *LiquidityController) LiquidityInfo() {
	liq.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", liq.Ctx.Request.Header.Get("Origin"))
	var response bo.ResponseBo
	response.Code = "00001"
	response.Message = "fail"

	chainType := liq.GetString("chainType")
	lpId, _ := liq.GetInt("lpId")
	var service service.LiquidityService
	response, err := service.LiquidityInfo(chainType, lpId)
	if err != nil {
		logs.Error("service.LiquidityInfo() is err: ", err)
	}
	liq.Data["json"] = response
	liq.ServeJSON()
}

// 更新流动池信息
func (liq *LiquidityController) UpdateLiquidityInfo() {
	liq.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", liq.Ctx.Request.Header.Get("Origin"))
	var response bo.ResponseBo
	response.Code = "00000"
	response.Message = "success"

	// 请求参数
	data := liq.Ctx.Input.RequestBody
	var liqBo bo.UpdateLiquidityInfoBo
	err := json.Unmarshal(data, &liqBo)
	if err != nil {
		logs.Error("json.Unmarshal(data, &liqBo) is err: ", err)
		return
	}
	logs.Info("UpdateLiquidityInfo请求参数：", liqBo)
	if !validChainType(liqBo.ChainType) {
		logs.Error("chainType is error")
		return
	}

	// 验证签名
	paramMap := make(map[string]interface{})
	paramMap["chainType"] = liqBo.ChainType
	paramMap["lpId"] = liqBo.LpId
	if ok, _ := util.VerifySign(paramMap, liqBo.Sign, "&"); !ok {
		response.Code = "00004"
		response.Message = "签名错误"
	}

	// 更新流动池信息
	task.UpdateLiquidityInfoPanic(liqBo.ChainType)
	liq.Data["json"] = response
	liq.ServeJSON()
}
