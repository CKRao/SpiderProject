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
	"time"
)

func MZiTuParser(url string, header *map[string]string, wg *sync.WaitGroup) {
	defer wg.Done()
	response, err := httpmodule.GetResponse(url, header, false)

	if err != nil {
		fmt.Println(err)
		return
	}

	document, err := goquery.NewDocumentFromReader(bytes.NewReader(response))

	if err != nil {
		fmt.Println(err)
		return
	}

	//获取最后的页数
	lastPageNode := document.Find("body > div.main > div.content > div.pagenavi > a:nth-child(7) > span")
	lastPageNum, err := strconv.Atoi(lastPageNode.Text())
	fmt.Println("最后一页页码为", lastPageNum)

	if lastPageNum <= 0 {
		return
	}

	if err != nil {
		log.Fatal(err)
		return
	}
	//以页面图集标题为目录名

	dirPath := "E:/downImg/" + document.Find(".main-title").Text()
	exists, err := PathExists(dirPath)

	if err != nil {
		fmt.Println(err)
		return
	}

	if !exists {
		err = os.Mkdir(dirPath, os.ModePerm)

		if err != nil {
			fmt.Println(err)
			return
		}
	}

	for start := 1; start <= lastPageNum; start++ {
		wg.Add(1)
		go downloadImg(start, header, wg, url, dirPath)
		time.Sleep(1 * time.Second)
	}

}

func downloadImg(start int, header *map[string]string, wg *sync.WaitGroup, url string, dirPath string) {
	defer wg.Done()
	fmt.Printf("开始获取第 %d 页数据\n", start)
	resp, err := httpmodule.GetResponse(url+strconv.Itoa(start), header, false)

	if err != nil {
		fmt.Println(err)
		return
	}

	document, err := goquery.NewDocumentFromReader(bytes.NewReader(resp))

	if err != nil {
		fmt.Println(err)
		return
	}

	imageNode := document.Find(".main-image > p > a > img")
	src, exists := imageNode.Attr("src")

	if exists {
		imgResp, err := httpmodule.GetResponse(src, header, false)
		if err != nil {
			fmt.Println(err)
			return
		}
		filemodule.CreateImage(dirPath, src, imgResp)
	} else {
		fmt.Println("图片地址未找到！")
	}
}

/**
判断文件是否存在
*/
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
