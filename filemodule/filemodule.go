package filemodule

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"spiderProject/util"
)

func CreateImage(dirPath string, imageSrc string, imgResp []byte) {
	fmt.Printf("开始创建图片 %s \n", imageSrc)

	fileName := path.Base(imageSrc)

	filePath := dirPath + "/" + fileName

	exists, err := util.PathExists(filePath)

	if err != nil {
		fmt.Println(err)
		return
	}

	if exists {
		fmt.Println(filePath, "已存在")
		return
	}

	file, _ := os.Create(filePath)
	writer := bufio.NewWriter(file)
	_, err = io.Copy(writer, bytes.NewReader(imgResp))

	if err != nil {
		fmt.Println(err)
		return
	}
}
