// Copyright (c) 2020 Higker
// Open Source: MIT License
// Author: higker <deen.job@qq.com>
package main

import (
	"fmt"
	"sync"
)

//go语言里面的sync锁
//锁就用来处理go里面的协程并发问题的
//程序在多个goroutine操作一个全局变量时需要加锁
//在用同一时间来操作一个全局变量 mutex 互斥锁

var (
	x = 0
	wg sync.WaitGroup //等待组
	lock sync.Mutex //互斥锁
)
func add(){
	defer wg.Done()
	for i:=0;i<500000 ;i++ {
		lock.Lock() //加锁
		 x = x + 1
		lock.Unlock() //解锁
	}
}
func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
