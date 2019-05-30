package filemodule

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

func CreateImage(dirPath string, imageSrc string, imgResp []byte) {
	fmt.Printf("开始创建图片 %s \n", imageSrc)

	fileName := path.Base(imageSrc)
	file, _ := os.Create(dirPath + "/" + fileName)
	writer := bufio.NewWriter(file)
	_, err := io.Copy(writer, bytes.NewReader(imgResp))

	if err != nil {
		log.Fatal(err)
	}
}
