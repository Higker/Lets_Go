package main

import (
	"fmt"
	"os"
)

const (
	size     = 1024
	filePath = "D:/Go_workspace/src/Lets_Go/lets_18_file/ks.txt"
)

//go语言普通方法打开文件
func main() {
	file, err := os.Open(filePath) //打开文件
	defer file.Close()             //一定要记得关闭文件
	if err != nil {
		fmt.Println("文本打开失败~", err)
	}
	content := make([]byte, size*10)
	temp := make([]byte, size) //用存储读取文件的内存的切片
	for {

		line, err := file.Read(temp) //读取文件
		if err != nil {
			fmt.Println("文件内容读取错误~")
		}
		fmt.Println("本次读取了:", line, "字节")
		content = append(content, temp...)
		if line != size {
			break
		}
	}
	fmt.Println(string(content))
	var stop string
	fmt.Scan(&stop)
}
