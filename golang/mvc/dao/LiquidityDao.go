package dao

import (
	"jamSwap/mvc/po"
)

type LiquidityDao struct {
}

func (dao *LiquidityDao) LiquidityList(chainType string) []po.LiquidityInfo {
	var liquidity []po.LiquidityInfo
	o.Raw("select id, lp_id, lp_address, is_lp, total_liquidity, lp_price, total_amount, apy, chain_type"+
		" from liquidity_info where chain_type = ? and is_delete = 0", chainType).QueryRows(&liquidity)
	return liquidity
}

func (dao *LiquidityDao) Update(id int, liquidityNew string, priceNew string, amountNew string, apyNew string) (int64, error) {
	liquidityInfo := po.LiquidityInfo{Id: id}
	liquidityInfo.TotalLiquidity = liquidityNew
	liquidityInfo.LpPrice = priceNew
	liquidityInfo.TotalAmount = amountNew
	liquidityInfo.Apy = apyNew
	return o.Update(&liquidityInfo, "TotalLiquidity", "LpPrice", "TotalAmount", "Apy")
}

func (dao *LiquidityDao) LiquidityInfo(chainType string, lpId int) po.LiquidityInfo {
	var liquidity po.LiquidityInfo
	o.Raw("select id, lp_id, lp_address, is_lp, total_liquidity, lp_price, total_amount, apy, chain_type"+
		" from liquidity_info where chain_type = ? and lp_id = ? and is_delete = 0", chainType, lpId).QueryRow(&liquidity)
	return liquidity
}
