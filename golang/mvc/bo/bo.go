package bo

import "jamSwap/mvc/po"

// 接口返回参数
type ResponseBo struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type UserAdvertisementBo struct {
	UserAddr  string `json:"userAddr"`
	ImgUrl    string `json:"imgUrl"`
	ChainType string `json:"chainType"`
}

type RemoveAdvertisementBo struct {
	UserAddr  string `json:"userAddr"`
	ChainType string `json:"chainType"`
	IsDelete  int    `json:"isDelete"`
}

type HasAdvertisementResponse struct {
	Has               bool `json:"has"`
	IsDelete          int  `json:"isDelete"`
	UserAdvertisement po.UserAdvertisement
}

type TotalTvlAndJamPriceResponse struct {
	TvlNum    string `json:"tvlNum"`
	TvlAmount string `json:"tvlAmount"`
	JamPrice  string `json:"jamPrice"`
}

type UpdateLiquidityInfoBo struct {
	ChainType string `json:"chainType"`
	LpId      int    `json:"lpId"`
}
