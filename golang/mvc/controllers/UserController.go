package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"jamSwap/mvc/bo"
	"jamSwap/mvc/service"
	"path"
	"strconv"
	"time"
)

type UserController struct {
	beego.Controller
}

// 保存图片
// @return 返回保存的图片URL
func (user *UserController) SaveFile() {
	user.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", user.Ctx.Request.Header.Get("Origin"))
	var response bo.ResponseBo
	response.Code = "00001"
	response.Message = "fail"

	// 保存图片
	file, _, err := user.GetFile("file")
	if err != nil {
		logs.Error("File retrieval failure: ", err)
	} else {
		defer file.Close()
		filePath := beego.AppConfig.String("home::filePath")
		databaseFilePath := beego.AppConfig.String("home::databaseFilePath")
		fn := strconv.Itoa(int(time.Now().UnixNano())) + ".png"
		fileName := path.Join(filePath, fn)
		fileName2 := path.Join(databaseFilePath, fn)
		logs.Info("fileName: ", fileName)
		err = user.SaveToFile("file", fileName)
		if err != nil {
			logs.Error("user.SaveToFile error: ", err)
		} else {
			response.Code = "00000"
			response.Message = "success"
			response.Data = fileName2
		}
	}

	user.Data["json"] = response
	user.ServeJSON()
}

// 保存用户的广告信息
// @param userAddr 用户地址
// @param imgUrl 广告图片URL
// @param chainType 链类型 ETH,BSC,HECO,DOT
func (user *UserController) SaveAdvertisement() {
	user.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", user.Ctx.Request.Header.Get("Origin"))
	var response bo.ResponseBo
	response.Code = "00001"
	response.Message = "fail"

	// 请求参数
	data := user.Ctx.Input.RequestBody
	var userBo bo.UserAdvertisementBo
	err := json.Unmarshal(data, &userBo)
	if err != nil {
		logs.Error("json.Unmarshal(data, &userBo) is err: ", err)
		return
	}
	logs.Info("SaveAdvertisement请求参数：", userBo)
	if !validChainType(userBo.ChainType) {
		logs.Error("chainType is error")
		return
	}

	// 保存用户的广告信息
	var service service.UserService
	response, err = service.SaveAdvertisement(userBo)
	if err != nil {
		logs.Error("service.SaveAdvertisement(userBo) is err: ", err)
	}
	user.Data["json"] = response
	user.ServeJSON()
}

// 判断用户是否已拥有广告，true有，false没有
// @param userAddr 用户地址
// @param chainType 链类型 ETH,BSC,HECO,DOT
// @return true有，false没有，如果返回true的话，会同时返回广告信息
func (user *UserController) HasAdvertisement() {
	user.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", user.Ctx.Request.Header.Get("Origin"))
	var response bo.ResponseBo
	response.Code = "00001"
	response.Message = "fail"

	// 请求参数
	userAddr := user.GetString("userAddr")
	chainType := user.GetString("chainType")
	logs.Info("HasAdvertisement请求参数：userAddr=", userAddr, "，chainType=", chainType)
	if !validChainType(chainType) {
		logs.Error("chainType is error")
		return
	}

	var service service.UserService
	response, err := service.HasAdvertisement(userAddr, chainType)
	if err != nil {
		logs.Error("service.HasAdvertisement(has) is err: ", err)
	}
	user.Data["json"] = response
	user.ServeJSON()
}

// 获取广告列表
// @param userAddr 用户地址，如果是管理员地址要做特殊处理
// @param chainType 链类型
// @param pageNo 第几页，从1开始
// @param pageSize 页大小，默认为8
// @return 广告列表
func (user *UserController) List() {
	user.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", user.Ctx.Request.Header.Get("Origin"))
	var response bo.ResponseBo
	response.Code = "00001"
	response.Message = "fail"

	userAddr := user.GetString("userAddr")
	chainType := user.GetString("chainType")
	pageNo, _ := user.GetInt("pageNo")
	pageSize, _ := user.GetInt("pageSize")
	if pageSize == 0 {
		pageSize = 8
	}

	// 获取广告列表
	var service service.UserService
	response, err := service.List(userAddr, chainType, pageNo, pageSize)
	if err != nil {
		logs.Error("service.List() is err: ", err)
	}
	user.Data["json"] = response
	user.ServeJSON()
}

// 获取所有地址
// @param chainType 链类型
// @return 地址列表
func (user *UserController) AddressList() {
	user.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", user.Ctx.Request.Header.Get("Origin"))
	var response bo.ResponseBo
	response.Code = "00001"
	response.Message = "fail"

	chainType := user.GetString("chainType")

	// 获取广告列表
	var service service.UserService
	addressList, err := service.AddressList(chainType)
	if err != nil {
		logs.Error("service.AddressList() is err: ", err)
	}
	user.Data["json"] = addressList
	user.ServeJSON()
}

// 移除广告
// @param userAddr 用户地址
// @param chainType 链类型 ETH,BSC,HECO,DOT
// @return 00000成功，其它失败
func (user *UserController) RemoveAdvertisement() {
	user.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", user.Ctx.Request.Header.Get("Origin"))
	var response bo.ResponseBo
	response.Code = "00001"
	response.Message = "fail"

	// 请求参数
	data := user.Ctx.Input.RequestBody
	var removeBo bo.RemoveAdvertisementBo
	err := json.Unmarshal(data, &removeBo)
	if err != nil {
		logs.Error("json.Unmarshal(data, &userBo) is err: ", err)
		return
	}
	logs.Info("RemoveAdvertisement请求参数：", removeBo)
	if !validChainType(removeBo.ChainType) {
		logs.Error("chainType is error")
		return
	}

	// 移除用户的广告信息
	var service service.UserService
	response, err = service.RemoveAdvertisement(removeBo)
	if err != nil {
		logs.Error("service.RemoveAdvertisement(removeBo) is err: ", err)
	}
	user.Data["json"] = response
	user.ServeJSON()
}

// 隐藏广告
// @param id 广告地址
// @return 00000成功，其它失败
func (user *UserController) HideAdvertisement() {
	user.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", user.Ctx.Request.Header.Get("Origin"))
	var response bo.ResponseBo
	response.Code = "00001"
	response.Message = "fail"

	// 请求参数
	data := user.Ctx.Input.RequestBody
	var hideBo bo.OprAdvertisementBo
	err := json.Unmarshal(data, &hideBo)
	if err != nil {
		logs.Error("json.Unmarshal(data, &hideBo) is err: ", err)
		return
	}
	logs.Info("HideAdvertisement请求参数：", hideBo)

	// 隐藏广告
	var service service.UserService
	response, err = service.HideAdvertisement(hideBo)
	if err != nil {
		logs.Error("service.HideAdvertisement(hideBo) is err: ", err)
	}
	user.Data["json"] = response
	user.ServeJSON()
}

// 开放广告
// @param id 广告地址
// @return 00000成功，其它失败
func (user *UserController) OpenAdvertisement() {
	user.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", user.Ctx.Request.Header.Get("Origin"))
	var response bo.ResponseBo
	response.Code = "00001"
	response.Message = "fail"

	// 请求参数
	data := user.Ctx.Input.RequestBody
	var openBo bo.OprAdvertisementBo
	err := json.Unmarshal(data, &openBo)
	if err != nil {
		logs.Error("json.Unmarshal(data, &openBo) is err: ", err)
		return
	}
	logs.Info("OpenAdvertisement请求参数：", openBo)

	// 隐藏广告
	var service service.UserService
	response, err = service.OpenAdvertisement(openBo)
	if err != nil {
		logs.Error("service.OpenAdvertisement(openBo) is err: ", err)
	}
	user.Data["json"] = response
	user.ServeJSON()
}

// 修改广告的排序序号
// @param id 广告地址
// @param sortNum
// @return 00000成功，其它失败
func (user *UserController) UpdateSortNum() {
	user.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", user.Ctx.Request.Header.Get("Origin"))
	var response bo.ResponseBo
	response.Code = "00001"
	response.Message = "fail"

	// 请求参数
	data := user.Ctx.Input.RequestBody
	var sortBo bo.UpdateSortNumBo
	err := json.Unmarshal(data, &sortBo)
	if err != nil {
		logs.Error("json.Unmarshal(data, &sortBo) is err: ", err)
		return
	}
	logs.Info("UpdateSortNum请求参数：", sortBo)

	// 修改广告的排序序号
	var service service.UserService
	response, err = service.UpdateSortNum(sortBo)
	if err != nil {
		logs.Error("service.UpdateSortNum(sortBo) is err: ", err)
	}
	user.Data["json"] = response
	user.ServeJSON()
}

func validChainType(chainType string) bool {
	if chainType != "ETH" && chainType != "BSC" && chainType != "HECO" && chainType != "DOT" {
		return false
	} else {
		return true
	}
}
