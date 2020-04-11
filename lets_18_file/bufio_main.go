package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//通过bufio读写文件
const (
	//size     = 1024
	filePath = "D:/Go_workspace/src/Lets_Go/lets_18_file/ks.txt"
)

func main() {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		fmt.Println("文件打开失败！！")
	}
	reader := bufio.NewReader(file)
	content := ""
	for {
		temp, err := reader.ReadString('\n') //替换每行末尾换行符
		content += temp                      //追加内容
		if err == io.EOF {
			fmt.Println("文件读取完毕")
			break
		}
	}
	//cannot convert content (type []string) to type string
	fmt.Println(content)
	var stop string
	fmt.Scan(&stop)
}
