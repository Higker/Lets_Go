// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/4/21 - 3:30 下午

//go 并发编程之goroutine
package main

import (
	"fmt"
)

func SayHi() {
	fmt.Println("Hello World.")
}

//go语言main函数启动之后程序会自动创建一个主goroutine去执行主函数
func main() {
	go SayHi() //通过go关键字去启动一个goroutine让它去执行我们的自己自定义的任务函数
	fmt.Println("您好世界~")
	//time.Sleep(time.Second * 2) //程序阻塞2秒 因为程序执行的太快了.
	for i := 1; i <= 100; i++ {
		go func(int) {
			fmt.Println(i) //这里打印的无序的因为for速度快  但是goroutine启动执行任务需要时间
		}(i)
	}
	for i := 1; i <= 100; i++ {
		go func(i int) { // 通过副本解决
			fmt.Println(i)
		}(i)
	}
}
