package main

import (
	"fmt"
	"math/rand"
	"time"
)

//go语言中的rand 生成随机数
func main() {
	fmt.Println(rand.Int())
	//0 到 N位的数
	fmt.Println(rand.Intn(100))
	//设置种子数
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Int())
	//生成指定范围的随机数
	fmt.Println(rand.Intn(10) + 1)
}
