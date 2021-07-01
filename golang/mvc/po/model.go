package po

import "time"

// 用户广告信息
type UserAdvertisement struct {
	Id         int       `json:"id"`
	UserAddr   string    `json:"userAddr"`
	ImgUrl     string    `json:"imgUrl"`    // 广告url
	ChainType  string    `json:"chainType"` // 链类型：ETH,HECO,BSC,DOT
	IsDelete   int       `json:"isDelete"`
	IsHide     int       `json:"isHide"`
	SortNum    int       `json:"sortNum"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}

// 流动池信息
type LiquidityInfo struct {
	Id             int    `json:"id"`
	LpId           int    `json:"lpId"` // 对应合约中的流动池id
	LpName         string `json:"lpName"`
	LpAddress      string `json:"lpAddress"`      // 流动性代币地址
	IsLp           int    `json:"isLp"`           // 0单币种，1双币种
	TotalLiquidity string `json:"totalLiquidity"` // 总流动性(流动性代币个数)
	LpPrice        string `json:"lpPrice"`        // 流动性代币的价格
	TotalAmount    string `json:"totalAmount"`    // 总流动性金额(所有流动性代币换算成USDT的数量)
	Apy            string `json:"apy"`            // 年化收益率
	ChainType      string `json:"chainType"`      // 链类型：ETH,HECO,BSC,DOT
	IsDelete       int    `json:"isDelete"`       // 0有效，1无效
}
type TConfig struct {
	Id        int    `json:"id"`
	TKey      string `json:"tKey"`
	TValue    string `json:"tValue"`
	ChainType string `json:"chainType"`
}
