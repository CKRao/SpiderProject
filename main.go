package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/puerkitobio/goquery"
	"io"
	"log"
	"os"
	"path"
	"spiderProject/httpmodule"
	"strconv"
)

const (
	BaseUrl = "https://www.mzitu.com/181701/"
)

func main() {
	header := make(map[string]string)

	header["Referer"] = "https://www.mzitu.com/"

	response, err := httpmodule.GetResponse(BaseUrl, &header, false)

	if err != nil {
		log.Fatal(err)
	}

	document, err := goquery.NewDocumentFromReader(bytes.NewReader(response))

	if err != nil {
		log.Fatal(err)
	}

	//获取最后的页数
	lastPageNode := document.Find("body > div.main > div.content > div.pagenavi > a:nth-child(7) > span")
	lastPageNum, err := strconv.Atoi(lastPageNode.Text())
	fmt.Println("最后一页页码为", lastPageNum)

	if err != nil {
		log.Fatal(err)
	}

	for start := 1; start <= lastPageNum; start++ {
		fmt.Printf("开始获取第 %d 页数据\n",start)
		response, err := httpmodule.GetResponse(BaseUrl + strconv.Itoa(start), &header, false)

		if err != nil {
			log.Fatal(err)
		}

		document, err := goquery.NewDocumentFromReader(bytes.NewReader(response))

		if err != nil {
			log.Fatal(err)
		}

		imageNode := document.Find(".main-image > p > a > img")
		src, exists := imageNode.Attr("src")

		if exists {
			imgResp, err := httpmodule.GetResponse(src, &header, false)
			if err != nil {
				log.Fatal(err)
			}
			fileName := path.Base(src)
			file, _ := os.Create("C:/Users/clarkrao/go/src/spiderProject/jpg7/" + fileName)
			writer := bufio.NewWriter(file)
			io.Copy(writer, bytes.NewReader(imgResp))
		} else {
			fmt.Println("图片地址未找到！")
		}
	}

}
