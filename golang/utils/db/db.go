package db

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"jamSwap/mvc/po"
)

func Init() {
	// 初始化Mysql
	url := beego.AppConfig.String("mysql::url")
	port := beego.AppConfig.String("mysql::port")
	dbName := beego.AppConfig.String("mysql::dbName")
	user := beego.AppConfig.String("mysql::user")
	password := beego.AppConfig.String("mysql::password")
	maxIdle, _ := beego.AppConfig.Int("mysql::maxIdle")
	maxConn, _ := beego.AppConfig.Int("mysql::maxConn")
	sqlConn := user + ":" + password + "@tcp(" + url + ":" + port + ")/" + dbName + "?charset=utf8"

	err := orm.RegisterDataBase("default", "mysql", sqlConn, maxIdle, maxConn)
	if err != nil {
		logs.Error("Mysql Init Failure！, err:", err)
		return
	}

	orm.RegisterModel(new(po.UserAdvertisement), new(po.LiquidityInfo), new(po.TConfig))
	orm.RunSyncdb("default", false, true)

	logs.Info("Mysql Init Successfully!")
}
