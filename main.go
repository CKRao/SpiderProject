package main

import (
	"bytes"
	"fmt"
	"github.com/puerkitobio/goquery"
	"log"
	"os"
	"spiderProject/httpmodule"
	"spiderProject/parsemoudule"
	"spiderProject/util"
	"strconv"
	"sync"
	"time"
)

const (
	BaseUrl = "https://www.mzitu.com/tag/meitun/page/"
)

func main() {
	wg := &sync.WaitGroup{}

	header := make(map[string]string)

	header["referer"] = "https://www.mzitu.com/"

	exits, err := util.PathExists("./img")

	if err != nil {
		log.Fatal(err)
	}

	if !exits {
		//创建图片存放目录
		err := os.Mkdir("./img", os.ModePerm)

		if err != nil {
			log.Fatal(err)
		}
	}

	for start := 1; start <= 20; start++ {
		wg.Add(1)
		go startFirstPage(header, wg, start)
		time.Sleep(20 * time.Second) //延迟30秒去处理下一个任务，不然可能造成响应数据拿不到
	}

	wg.Wait()
	fmt.Println("ALL DOWN")
}

func startFirstPage(header map[string]string, wg *sync.WaitGroup, index int) {
	defer wg.Done()
	response, err := httpmodule.GetResponse(BaseUrl+strconv.Itoa(index), &header, false)

	if err != nil {
		fmt.Println(err)
		return
	}

	document, err := goquery.NewDocumentFromReader(bytes.NewReader(response))

	if err != nil {
		fmt.Println(err)
		return
	}

	urlPath := make([]string, 24)
	document.Find("#pins > li > a").Each(func(i int, selection *goquery.Selection) {
		val, exists := selection.Attr("href")
		if exists && len(val) > 0 && val != "" {
			urlPath = append(urlPath, val+"/")
		}
	})

	for index, url := range urlPath {
		if url != "" {
			wg.Add(1)
			header["referer"] = url
			go parsemoudule.MZiTuParser(url, &header, wg)
			if index%2 == 0 {
				time.Sleep(10 * time.Second) //延迟10秒去处理下一个任务，不然可能造成响应数据拿不到
				//runtime.Gosched()
			}
		}
	}
}
