package main

import (
	"bytes"
	"fmt"
	"github.com/puerkitobio/goquery"
	"log"
	"os"
	"spiderProject/httpmodule"
	"spiderProject/parsemoudule"
)

const (
	BaseUrl = "https://www.mzitu.com/"
)

func main() {
	//wg := &sync.WaitGroup{}

	header := make(map[string]string)

	header["referer"] = "https://www.mzitu.com/"

	exits, err := PathExists("./img")

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

	response, err := httpmodule.GetResponse(BaseUrl, &header, false)

	if err != nil {
		log.Fatal(err)
	}

	document, err := goquery.NewDocumentFromReader(bytes.NewReader(response))

	if err != nil {
		log.Fatal(err)
	}

	urlPath := make([]string, 24)
	document.Find("#pins > li > a").Each(func(i int, selection *goquery.Selection) {
		val, exists := selection.Attr("href")
		if exists && len(val) > 0 && val != "" {
			urlPath = append(urlPath, val+"/")
		}
	})

	for _, url := range urlPath {
		if url != "" {
			//wg.Add(1)
			header["referer"] = url
			parsemoudule.MZiTuParser(url, &header, nil)
		}
	}

	//wg.Wait()
	fmt.Println("ALL DOWN")
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
