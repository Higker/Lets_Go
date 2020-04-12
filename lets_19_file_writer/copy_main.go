package main

import (
	"fmt"
	"io"
	"os"
)

const (
	filePath = "D:/Go_workspace/src/Lets_Go/lets_19_file_writer/test.txt"
	copyPath = "D:/Go_workspace/src/Lets_Go/lets_19_file_writer/test_copy.txt"
)

//go拷贝文件
func main() {
	io.Copy(dstFile(copyPath), srcFile(filePath))
	fmt.Println("copy file succeed.")
}

//目标文件
func dstFile(dstPath string) *os.File {
	file, err := os.OpenFile(dstPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("写出文件出错~")
	}
	return file
}

//源文件
func srcFile(srcPath string) *os.File {
	file, err := os.Open(srcPath)
	if err != nil {
		fmt.Println("读取源文件出错~")
	}
	return file
}
