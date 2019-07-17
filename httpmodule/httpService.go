package httpmodule

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"net/url"
	"spiderProject/ippool"
	"strings"
)

/**
Get请求获取数据
*/
func GetResponse(requrl string, headers *map[string]string, ok bool) ([]byte, error) {
	request, err := http.NewRequest("GET", requrl, nil)

	if err != nil {
		return nil, fmt.Errorf("create request error : %s", err)
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36")
	//循环设置请求头
	for key, value := range *headers {
		request.Header.Add(key, value)
	}

	client := &http.Client{}

	ip := ippool.GetIP()
	if ip == "" {
		client = http.DefaultClient
	} else {
		proxy, err := url.Parse(ip)
		client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxy),
			},
		}
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	//fmt.Println("Status : ", response.Status)
	if response.StatusCode >= 300 && response.StatusCode <= 500 {
		return nil, err
	}

	if ok {
		utf8Content := transform.NewReader(response.Body, simplifiedchinese.GBK.NewDecoder())
		return ioutil.ReadAll(utf8Content)
	}

	bytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return bytes, nil
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
		return nil, err
	}

	defer response.Body.Close()

	fmt.Println("StatusCode : ", response.StatusCode)

	if response.StatusCode >= 300 && response.StatusCode <= 500 {
		return nil, err
	}

	if ok {
		utf8Content := transform.NewReader(response.Body, simplifiedchinese.GBK.NewDecoder())
		return ioutil.ReadAll(utf8Content)
	}

	return ioutil.ReadAll(response.Body)
}
