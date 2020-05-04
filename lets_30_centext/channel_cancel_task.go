package main

import (
	"fmt"
	"time"
)

var (
	//初始化一个channel作为通知器
	stop = make(chan struct{})
)

func main() {
	go func() {
		for {
			select {
			//如果stop的channel能取出来值就说明要cancel任务
			case <-stop:
				fmt.Println("task stop...")
				return
			default:
				task()
			}
		}
	}()
	fmt.Println(stop)
	time.Sleep(5 * time.Second)
	stop <- struct{}{}
}

func task() {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("task running....")
}
