// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/4/21 - 6:09 下午

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	//这个是go里面用来配置goroutine使用的
	//只能操作指针 因为他是一个struct 值类型！！！
	wg sync.WaitGroup
	//存放我们for的数字方便观察
	sls []int
)

//go语言中的多个goroutine
//WaitGroup
func main() {
	fmt.Println("goroutine begin:", sls)
	for i := 0; i <= 10; i++ {
		// fatal error: all goroutines are asleep - deadlock!
		//注意这里加的1不是i如果是i的话每次都是加的不一样然后出现👆的异常！！！
		wg.Add(1)
		go task(i)
	}

	fmt.Println("InService:", sls)
	wg.Wait()
	fmt.Println("End Over:", sls)
}

// fatal error: all goroutines are asleep - deadlock!
func task(num int) {
	defer func() {
		wg.Done()
	}()
	//随机休眠几毫秒
	time.Sleep(duration())
	//将for循环的i存入到切片中我们等下好观察,
	//打印控制台太慢了回漏掉
	sls = append(sls, num) //不安全
}

//生成一个随机的等待时间 = 毫秒
func duration() time.Duration {
	rand.Seed(time.Now().UnixNano())
	return time.Millisecond * time.Duration(rand.Intn(3000))
}
