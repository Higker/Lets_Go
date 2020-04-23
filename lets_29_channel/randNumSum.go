// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/4/23 - 4:27 下午

package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
使用goroutine和channel实现个计 算int64随机数各 位数和的程序。
1.开启一个goroutine循环生成int64类型的随机数， 发送到jobChan
2.开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
3.主goroutine从resultChan取出结果并打印到终端输出
*/
func main() {
	rsChan := make(chan int64, 20)
	jobChan := make(chan int64, 10)
	go randInt64Num(jobChan)

	go computeSum(jobChan, rsChan)

	for e := range jobChan {
		fmt.Println(e)
	}
}

func randInt64Num(job chan<- int64) {
	rand.Seed(time.Now().UnixNano())
	job <- rand.Int63()
}
func computeSum(jobChan <-chan int64, result chan<- int64) {
	num := <-jobChan
	fmt.Println(num)
	var sum int64 = 0
	var i int64 = 10 // 初始化一个10 每次循环就是自己的平方
	for num >= 0 {
		sum += num % i // 123 -> 3 | 12 -> 2 | 10 -> 1
		i *= 10
	}
	result <- sum
}
