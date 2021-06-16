package httpUtils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/nahid/gohttp"
	"io/ioutil"
	"net/http"
)

func AsyncPost(url string, param map[string]interface{}) {
	req := gohttp.NewRequest()
	ch := make(chan *gohttp.AsyncResponse)
	// req.FormData(param).AsyncPost(url, ch)
	req.JSON(param).AsyncPost(url, ch)
}

func HttpPost(url string, params map[string]interface{}) (string, error) {
	data, err := json.Marshal(params)
	if err != nil {
		logs.Error(err.Error())
		return "", err
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		logs.Error(err.Error())
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logs.Error(err.Error())
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("请求失败")
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error(err.Error())
		return "", err
	}
	return string(respBytes), nil
}

func HttpGet(url string) ([]byte, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logs.Error(err.Error())
		return nil, err
	}
	apiKey := beego.AppConfig.String("tron::apiKey")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36")
	request.Header.Set("TRON-PRO-API-KEY", apiKey)
	resp, err := client.Do(request)
	if err != nil {
		logs.Error(err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("请求失败")
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error(err.Error())
		return nil, err
	}
	return respBytes, nil
}
