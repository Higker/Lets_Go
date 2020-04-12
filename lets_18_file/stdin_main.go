package main

import (
	"bufio"
	"fmt"
	"os"
)

//通过os.stdin获取用户输入
func main() {
	fmt.Println("请输入内容:")
	rd := bufio.NewReader(os.Stdin)
	str, _ := rd.ReadString('\n')
	fmt.Println("你输入的是:", str)
}
