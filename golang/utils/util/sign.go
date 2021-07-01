package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"sort"
	"strings"
)

// 验证签名
func VerifySign(data map[string]interface{}, sign, sep string) (bool, error) {
	isVerifySign, _ := beego.AppConfig.Bool("isVerifySign")
	if isVerifySign {
		if strings.EqualFold(sign, "") {
			return false, nil
		}
		key := beego.AppConfig.String("key")
		str, err := SignParam(data, sep, key)
		return str == sign, err
	} else {
		// 测试环境不验证签名
		return true, nil
	}
}

// 参数签名
// 将参数按照ascii码升序排序，如果参数值为空不参与排序，排序好后再在最后加入"&key=key"，
// 再用md5进行一次hash得到byte数组，再将byte数组转成16进制字符串
func SignParam(data map[string]interface{}, sep string, key string) (string, error) {
	str, err := SortParam(data, sep)
	str += "&key=" + key
	logs.Info("排序后的参数:", str)
	// md5
	md5Sum := md5.Sum([]byte(str))
	md5Str := hex.EncodeToString(md5Sum[:])
	logs.Info("排序后的参数进行md5后:", md5Str)
	return md5Str, err
}

// 将参数按照ascii码升序排序
func SortParam(data map[string]interface{}, sep string) (string, error) {
	var list []string
	for k := range data {
		value := data[k]
		list = append(list, fmt.Sprintf("%s=%s", k, value))
	}
	sort.Strings(list)
	str := strings.Join(list, sep)
	return str, nil
}
