package httpmodule

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"strings"
)

/**
Get请求获取数据
*/
func GetResponse(url string, headers *map[string]string, ok bool) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, fmt.Errorf("create request error : %s", err)
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36")
	//循环设置请求头
	for key, value := range *headers {
		request.Header.Add(key, value)
	}

	client := http.DefaultClient

	response, err := client.Do(request)

	if err != nil {
		return nil, fmt.Errorf("get resonse error : %s", err)
	}

	defer response.Body.Close()

	fmt.Println("Status : ", response.Status)
	if response.StatusCode >= 300 && response.StatusCode <= 500 {
		return nil, fmt.Errorf(" StatusCode error : %d", response.StatusCode)
	}

	if ok {
		utf8Content := transform.NewReader(response.Body, simplifiedchinese.GBK.NewDecoder())
		return ioutil.ReadAll(utf8Content)
	}

	return ioutil.ReadAll(response.Body)
}

/**
POST请求获取数据
*/
func PostResponse(url string, body string, headers *map[string]string, ok bool) ([]byte, error) {
	payLoad := strings.NewReader(body)

	request, err := http.NewRequest("POST", url, payLoad)

	if err != nil {
		return nil, fmt.Errorf("create request error : %s", err)
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36")

	//循环设置请求头
	for key, value := range *headers {
		request.Header.Add(key, value)
	}

	client := http.DefaultClient

	response, err := client.Do(request)

	if err != nil {
		return nil, fmt.Errorf("get resonse error : %s", err)
	}

	defer response.Body.Close()

	fmt.Println("StatusCode : ", response.StatusCode)

	if response.StatusCode >= 300 && response.StatusCode <= 500 {
		return nil, fmt.Errorf(" StatusCode error : %d", response.StatusCode)
	}

	if ok {
		utf8Content := transform.NewReader(response.Body, simplifiedchinese.GBK.NewDecoder())
		return ioutil.ReadAll(utf8Content)
	}

	return ioutil.ReadAll(response.Body)
}
