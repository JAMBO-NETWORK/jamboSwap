package bo

import "jamSwap/mvc/po"

// 请求参数Base
type ReqBaseBo struct {
	Sign string `json:"sign"` // 请求参数的hash
}

// 接口返回参数
type ResponseBo struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type UserAdvertisementBo struct {
	ReqBaseBo
	UserAddr  string `json:"userAddr"`
	ImgUrl    string `json:"imgUrl"`
	ChainType string `json:"chainType"`
}

type RemoveAdvertisementBo struct {
	ReqBaseBo
	UserAddr  string `json:"userAddr"`
	ChainType string `json:"chainType"`
	IsDelete  int    `json:"isDelete"`
}

// 隐藏广告请求参数
type OprAdvertisementBo struct {
	ReqBaseBo
	Id int `json:"id"`
}

// 修改广告排序请求参数
type UpdateSortNumBo struct {
	ReqBaseBo
	Id      int `json:"id"`
	SortNum int `json:"sortNum"`
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
	ReqBaseBo
	ChainType string `json:"chainType"`
	LpId      int    `json:"lpId"`
}
