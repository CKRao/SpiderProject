package parsemoudule

import (
	"bytes"
	"fmt"
	"github.com/puerkitobio/goquery"
	"log"
	"os"
	"spiderProject/filemodule"
	"spiderProject/httpmodule"
	"strconv"
	"sync"
)

func MZiTuParser(url string, header *map[string]string, wg *sync.WaitGroup) {
	//defer wg.Done()
	response, err := httpmodule.GetResponse(url, header, false)

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
	//以页面图集标题为目录名
	dirPath := "./img/" + document.Find(".main-title").Text()

	err = os.Mkdir(dirPath, os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}

	for start := 1; start <= lastPageNum; start++ {
		fmt.Printf("开始获取第 %d 页数据\n", start)
		response, err := httpmodule.GetResponse(url+strconv.Itoa(start), header, false)

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
			imgResp, err := httpmodule.GetResponse(src, header, false)
			if err != nil {
				log.Fatal(err)
			}
			filemodule.CreateImage(dirPath, src, imgResp)
		} else {
			fmt.Println("图片地址未找到！")
		}
	}

}
