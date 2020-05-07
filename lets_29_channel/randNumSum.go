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
type Job struct {
	Number int64
}
type JobChan struct {
	Id  int64
	Job *Job
}
type Resutl struct {
	Jc  *JobChan
	sum int64
}

func main() {
	chanJob := make(chan *JobChan, 20)
	reChan := make(chan *Resutl, 20)
	go dataService(chanJob)
	for i := 0; i < 2; i++ {
		go consumeData(chanJob, reChan)
	}
	for e := range reChan {
		fmt.Printf("srcNum: %d -> Sum: %d\n", e.Jc.Id, e.sum)
	}
}

//模拟生产源源不断的数据
func dataService(jobChan chan<- *JobChan) {
	for {
		//创建一个job结构体
		job := &Job{Number: randInt64()}
		//发送到jobChan处理
		jobChan <- &JobChan{Id: job.Number, Job: job} //这里是方便后面观察使用
		Sleep()
	}
}

//消费数据服务
func consumeData(jobChan <-chan *JobChan, resultChan chan<- *Resutl) {
	for {
		job := <-jobChan      //获取到工作管道中的job结构体
		num := job.Job.Number //得到数据数字
		sum := int64(0)
		for num > 0 {
			sum += num % 10 // 得到个位数的数字
			num = num / 10  //向左移动小数点一位 方便下次继续获取到个位数
		}
		//将计算结果发送到result通道中
		rc := &Resutl{Jc: job, sum: sum}
		resultChan <- rc
	}
}

//随机生成一个int64数据
func randInt64() int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63()
}

//随机休眠几毫秒
func Sleep() time.Duration {
	return time.Millisecond * time.Duration(randInt64())
}
