package dao

import (
	"github.com/astaxie/beego/orm"
	"jamSwap/mvc/po"
)

type ConfigDao struct {
}

func (dao *ConfigDao) QueryConfigByKey(key, chainType string) string {
	var config po.TConfig
	o.Raw("select t_value from t_config where t_key = ? and chain_type = ?", key, chainType).QueryRow(&config)
	return config.TValue
}

func (dao *ConfigDao) UpdateConfig(key, value, chainType string) (int64, error) {
	return o.QueryTable("t_config").Filter("t_key", key).Filter("chain_type", chainType).Update(orm.Params{
		"t_value": value,
	})
}
