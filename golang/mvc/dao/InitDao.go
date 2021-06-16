package dao

import "github.com/astaxie/beego/orm"

var o orm.Ormer

func Init() {
	o = orm.NewOrm()
}
