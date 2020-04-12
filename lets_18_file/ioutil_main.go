package main

import (
	"fmt"
	"io/ioutil"
)

const (
	Path = "D:/Go_workspace/src/Lets_Go/lets_18_file/test.txt"
)

//通过ioutil读取文件内容
func main() {
	cont, err := ioutil.ReadFile(Path)
	if err != nil {
		fmt.Println("文件获取失败~")
	}
	fmt.Println(string(cont))
}
