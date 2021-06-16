package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"jamSwap/mvc/dao"
	_ "jamSwap/routers"
	"jamSwap/utils/db"
	"jamSwap/utils/log"
	"jamSwap/utils/task"
	"path/filepath"
)

func main() {
	// 多环境配置文件加载
	env := beego.AppConfig.String("runmode")
	logs.Info("adapterName:", env)
	configFilePath, err := filepath.Abs("conf/" + env + ".conf")
	if err != nil {
		logs.Error("configFilePathError", err)
	}
	err = beego.LoadAppConfig("ini", configFilePath)
	if err != nil {
		logs.Error("loadConfigFileError", err)
	}

	log.Init()
	db.Init()
	dao.Init()
	task.Init()
	beego.Run()
}
