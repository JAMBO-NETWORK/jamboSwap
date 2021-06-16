package log

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"time"
)

func Init() {
	adapter := beego.AppConfig.String("log::adapter")
	enableFuncCallDepth, _ := beego.AppConfig.Bool("log::enableFuncCallDepth")
	logLevel, _ := beego.AppConfig.Int("log::logLevel")
	logPath := beego.AppConfig.String("log::logPath")
	maxLines, _ := beego.AppConfig.Int("log::maxLines")
	maxSize, _ := beego.AppConfig.Int("log::maxSize")
	logs.SetLogger(adapter)
	logs.EnableFuncCallDepth(enableFuncCallDepth)
	jsonConfig := make(map[string]interface{})

	// fmt.Println("aa:",time.Now().Format("2006-01-02 15:04:05"))
	date := "log_" + time.Now().Format("2006-01-02")
	fileName := date + ".log"
	jsonConfig["filename"] = logPath + fileName
	jsonConfig["maxlines"] = maxLines
	jsonConfig["maxsize"] = maxSize
	jsonConfig["level"] = logLevel
	confStr, _ := json.Marshal(jsonConfig)

	logs.SetLogger(logs.AdapterFile, string(confStr))
	// 日志异步输出
	logs.Async()
	// 异步输出允许设置缓冲 chan 的大小
	logs.Async(1e3)

	logs.Info("Log Init Successfully!")
}
