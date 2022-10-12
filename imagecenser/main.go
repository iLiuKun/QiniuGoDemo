package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"github.com/qiniu/go-sdk/v7/auth"
)

func imgCensor() {
	body := make(map[string]interface{})

	data := make(map[string]interface{})

	data["uri"] = "https://qiniu.shandianw.cn/qiniu/voting/upload/2022/4/24/player/900wlg9g1sx0pj.png"
	params := make(map[string]interface{})
	params["scenes"] = [...]string{"pulp"}

	body["data"] = data
	body["params"] = params

	bytesData, err := json.Marshal(body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(bytesData))

	reader := bytes.NewReader(bytesData)

	url := "http://ai.qiniuapi.com/v3/image/censor"
	request, err := http.NewRequest("POST", url, reader)
	request.Header.Set("Content-Type", "application/json")

	defer request.Body.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	accessToken, signErr := mac.SignRequestV2(request)
	if signErr != nil {
		err = signErr
		return
	}

	PostSend(url, accessToken, bytesData, "application/json")
}

func PostSend(url string, acc string, bodyBytes []byte, contentType string) {
	reader := bytes.NewReader(bodyBytes)

	request, err := http.NewRequest("POST", url, reader)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Qiniu "+acc)
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("request请求异常", err.Error())
		return
	}

	defer resp.Body.Close()
	m := make(map[string]interface{})
	json.NewDecoder(resp.Body).Decode(&m)
	fmt.Println("resposne响应", m)
	fmt.Println("ceshi", resp.Status)
}

var (
	accessKey = os.Getenv("QINIU_ACCESS_KEY")
	secretKey = os.Getenv("QINIU_SECRET_KEY")
	mac       = auth.New(accessKey, secretKey)
)

func main() {
	imgCensor()
}
