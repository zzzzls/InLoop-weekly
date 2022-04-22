package main

/*
格式化 Notion 导出的周刊文件
*/

import (
	"fmt"
	"net/url"

	"bufio"
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var page string
	flag.StringVar(&page, "p", "0", "周刊期号")
	flag.Parse()

	dir, _ := os.Getwd()
	artilePath := filepath.Join(dir, "..", "Article", page)

	ImagePath, MarkDownFile := LookForFile(artilePath)
	NewImagePath := "image"

	fullMarkDownFile := filepath.Join(artilePath, MarkDownFile)
	fileContentByte, err := ioutil.ReadFile(fullMarkDownFile)
	if err != nil {
		panic(err)
	}

	fmt.Println(url.PathEscape(ImagePath))
	// 从文件首行读出文件名
	firstLine, _ := bufio.NewReader(strings.NewReader(string(fileContentByte))).ReadString('\n')
	NewMarkDownFile := strings.Split(firstLine, "：")[1]
	NewMarkDownFile = strings.TrimSpace(NewMarkDownFile) + ".md"

	fileContent := strings.ReplaceAll(string(fileContentByte), url.PathEscape(ImagePath), NewImagePath)
	ioutil.WriteFile(fullMarkDownFile, []byte(fileContent), 0666)

	// 重命名文件
	Rename(MarkDownFile, NewMarkDownFile, artilePath)
	Rename(ImagePath, NewImagePath, artilePath)
}

func LookForFile(artilePath string) (ImagePath string, MarkDownFile string) {

	folders, err := ioutil.ReadDir(artilePath)
	if err != nil {
		panic(err)
	}

	for _, item := range folders {
		filename := item.Name()

		if item.IsDir() {
			ImagePath = filename
			continue
		}
		if ok := strings.HasSuffix(filename, ".md"); ok {
			MarkDownFile = filename
			continue
		}

	}

	return ImagePath, MarkDownFile
}

func Rename(old, new string, rootpath string) {
	err := os.Rename(filepath.Join(rootpath, old), filepath.Join(rootpath, new))
	if err != nil {
		panic(err)
	}
}
