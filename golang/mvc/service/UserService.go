package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/ethereum/go-ethereum/log"
	"jamSwap/contracts/bsc/mortgage"
	"jamSwap/mvc/bo"
	"jamSwap/mvc/dao"
	"jamSwap/mvc/po"
	"jamSwap/utils/util"
	"math/big"
	"strings"
)

var base18 int64 = 1000000000000000000

type UserService struct {
}

func (service *UserService) SaveAdvertisement(userBo bo.UserAdvertisementBo) (bo.ResponseBo, error) {
	var response bo.ResponseBo
	// 验证签名
	paramMap := make(map[string]interface{})
	paramMap["userAddr"] = userBo.UserAddr
	paramMap["imgUrl"] = userBo.ImgUrl
	paramMap["chainType"] = userBo.ChainType
	if ok, err := util.VerifySign(paramMap, userBo.Sign, "&"); !ok {
		response.Code = "00004"
		response.Message = "签名错误"
		return response, err
	}

	response.Code = "00001"
	// 查询用户的链上信息，余额是否大于mortgageAmount
	mortgageAmount, _ := beego.AppConfig.Int("home::mortgageAmount")
	chainType := userBo.ChainType
	clientUrl := beego.AppConfig.String(chainType + "::clientUrl")
	mortgageAddress := beego.AppConfig.String(chainType + "::mortgageAddress")
	mortgageInstance, err := mortgage.GetMortgageInstance(clientUrl, mortgageAddress)
	if err != nil {
		logs.Error("GetMortgageInstance error: ", err)
		return response, err
	}
	chainMortgageAmountBig, err := mortgage.GetMortgageAmount(mortgageInstance, userBo.UserAddr)
	chainMortgageAmountBig.Div(chainMortgageAmountBig, big.NewInt(base18))
	chainMortgageAmount := int(chainMortgageAmountBig.Int64())
	if chainMortgageAmount < mortgageAmount {
		response.Message = "需要先质押才能创建广告"
		return response, errors.New("需要先质押才能创建广告")
	}

	// 每个地址只能同时拥有一个有效广告
	var dao dao.UserDao
	var userPo po.UserAdvertisement
	dataByte, _ := json.Marshal(userBo)
	json.Unmarshal(dataByte, &userPo)
	maxSortNum := dao.MaxSortNum()
	if maxSortNum == 0 {
		maxSortNum = 2
	} else {
		maxSortNum++
	}
	userPo.SortNum = maxSortNum
	user := dao.QueryAdvertisement(userBo.UserAddr, userBo.ChainType)
	if user.Id > 0 {
		if user.IsDelete == 0 {
			response.Message = "每个地址只能同时拥有一个有效广告"
			return response, errors.New("每个地址只能同时拥有一个有效广告")
		} else {
			// 修改
			userPo.Id = user.Id
			_, err = dao.UpdateUser(&userPo)
			if err == nil {
				response.Code = "00000"
				response.Message = "success"
			}
		}
	} else {
		// 没有新增
		_, err = dao.InsertUser(&userPo)
		if err == nil {
			response.Code = "00000"
			response.Message = "success"
		}
	}

	return response, nil
}

func (service *UserService) HasAdvertisement(userAddr, chainType string) (bo.ResponseBo, error) {
	var response bo.ResponseBo
	response.Code = "00000"
	response.Message = "success"
	var bo bo.HasAdvertisementResponse
	var dao dao.UserDao
	user := dao.QueryAdvertisement(userAddr, chainType)
	fmt.Println("user: ", user)
	if user.Id > 0 && user.IsDelete == 0 {
		bo.Has = true
		bo.UserAdvertisement = user
	} else {
		bo.Has = false
		bo.IsDelete = user.IsDelete
	}
	response.Data = bo
	return response, nil
}

// 获取广告列表
func (service *UserService) List(userAddr, chainType string, pageNo, pageSize int) (bo.ResponseBo, error) {
	var response bo.ResponseBo
	response.Code = "00000"
	response.Message = "success"
	var dao dao.UserDao
	totalCount := 0
	start := (pageNo - 1) * pageSize
	admin := beego.AppConfig.String("admin")
	var list *[]po.UserAdvertisement
	isOwner := 0
	if strings.EqualFold(admin, userAddr) {
		// 管理员地址，需要查询被隐藏的
		isOwner = 1
		totalCount = dao.TotalCount2(chainType)
		list = dao.List2(chainType, start, pageSize)
	} else {
		totalCount = dao.TotalCount(chainType)
		list = dao.List(chainType, start, pageSize)
	}
	page := bo.PageUtil(totalCount, pageNo, pageSize, list, isOwner)
	response.Data = page
	return response, nil
}

func (service *UserService) AddressList(chainType string) ([]string, error) {
	var dao dao.UserDao
	list := dao.AddressList(chainType)
	addressList := make([]string, len(list), len(list))
	for _, v := range list {
		addressList = append(addressList, v.UserAddr)
	}
	return addressList, nil
}

// 移除用户的广告信息
func (service *UserService) RemoveAdvertisement(removeBo bo.RemoveAdvertisementBo) (bo.ResponseBo, error) {
	var response bo.ResponseBo
	// 验证签名
	paramMap := make(map[string]interface{})
	paramMap["userAddr"] = removeBo.UserAddr
	paramMap["chainType"] = removeBo.ChainType
	paramMap["isDelete"] = removeBo.IsDelete
	if ok, err := util.VerifySign(paramMap, removeBo.Sign, "&"); !ok {
		response.Code = "00004"
		response.Message = "签名错误"
		return response, err
	}

	response.Code = "00001"
	response.Message = "fail"
	var dao dao.UserDao
	_, err := dao.UpdateAdvertisement(removeBo.UserAddr, removeBo.ChainType, removeBo.IsDelete)
	if err != nil {
		log.Error("UpdateAdvertisement is error: ", err)
	} else {
		response.Code = "00000"
		response.Message = "success"
	}
	return response, nil
}

// 隐藏广告
func (service *UserService) HideAdvertisement(hideBo bo.OprAdvertisementBo) (bo.ResponseBo, error) {
	var response bo.ResponseBo
	// 验证签名
	paramMap := make(map[string]interface{})
	paramMap["id"] = hideBo.Id
	if ok, err := util.VerifySign(paramMap, hideBo.Sign, "&"); !ok {
		response.Code = "00004"
		response.Message = "签名错误"
		return response, err
	}

	response.Code = "00001"
	response.Message = "fail"
	var dao dao.UserDao
	_, err := dao.HideAdvertisement(hideBo.Id)
	if err != nil {
		log.Error("HideAdvertisement is error: ", err)
	} else {
		response.Code = "00000"
		response.Message = "success"
	}
	return response, nil
}

// 开放广告
func (service *UserService) OpenAdvertisement(openBo bo.OprAdvertisementBo) (bo.ResponseBo, error) {
	var response bo.ResponseBo
	// 验证签名
	paramMap := make(map[string]interface{})
	paramMap["id"] = openBo.Id
	if ok, err := util.VerifySign(paramMap, openBo.Sign, "&"); !ok {
		response.Code = "00004"
		response.Message = "签名错误"
		return response, err
	}

	response.Code = "00001"
	response.Message = "fail"
	var dao dao.UserDao
	_, err := dao.OpenAdvertisement(openBo.Id)
	if err != nil {
		log.Error("OpenAdvertisement is error: ", err)
	} else {
		response.Code = "00000"
		response.Message = "success"
	}
	return response, nil
}

// 修改广告的排序序号
func (service *UserService) UpdateSortNum(sortBo bo.UpdateSortNumBo) (bo.ResponseBo, error) {
	var response bo.ResponseBo
	// 验证签名
	paramMap := make(map[string]interface{})
	paramMap["id"] = sortBo.Id
	paramMap["sortNum"] = sortBo.SortNum
	if ok, err := util.VerifySign(paramMap, sortBo.Sign, "&"); !ok {
		response.Code = "00004"
		response.Message = "签名错误"
		return response, err
	}

	response.Code = "00001"
	response.Message = "fail"
	var dao dao.UserDao
	_, err := dao.UpdateSortNum(sortBo.Id, sortBo.SortNum)
	if err != nil {
		log.Error("UpdateSortNum is error: ", err)
	} else {
		response.Code = "00000"
		response.Message = "success"
	}
	return response, nil
}
