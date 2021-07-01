package service

import (
	"jamSwap/mvc/bo"
	"jamSwap/mvc/dao"
	"strconv"
)

type LiquidityService struct {
}

// 获取所有流动池的总流动性和jamPrice
func (service *LiquidityService) GetTotalTvlAndJamPrice(chainType string) (bo.ResponseBo, error) {
	var dao dao.ConfigDao
	tvlNum := dao.QueryConfigByKey("tvlNum", chainType)
	tvlAmount := dao.QueryConfigByKey("tvlAmount", chainType)
	jamPrice := dao.QueryConfigByKey("jamPrice", chainType)

	totalTvlAndJamPriceResponse := bo.TotalTvlAndJamPriceResponse{
		TvlNum:    tvlNum,
		TvlAmount: tvlAmount,
		JamPrice:  jamPrice,
	}
	var response bo.ResponseBo
	response.Code = "00000"
	response.Message = "success"
	response.Data = totalTvlAndJamPriceResponse
	return response, nil
}

// 获取流动池信息
func (service *LiquidityService) LiquidityList(chainType string) (bo.ResponseBo, error) {
	var dao dao.LiquidityDao
	list := dao.LiquidityList(chainType)
	var response bo.ResponseBo
	response.Code = "00000"
	response.Message = "success"
	response.Data = list
	return response, nil
}

// 获取流动池信息
func (service *LiquidityService) LiquidityInfo(chainType string, lpId int) (bo.ResponseBo, error) {
	var dao dao.LiquidityDao
	liq := dao.LiquidityInfo(chainType, lpId)
	var response bo.ResponseBo
	if liq.Id > 0 {
		response.Code = "00000"
		response.Message = "success"
		response.Data = liq
	} else {
		response.Code = "00001"
		response.Message = "lpId=" + strconv.Itoa(lpId) + "不存在"
	}
	return response, nil
}
