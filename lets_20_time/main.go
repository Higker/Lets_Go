package main

import (
	"fmt"
	"time"
)

//go里面的时间操作
func main() {
	d := time.Now()
	fmt.Println("年:", d.Year())
	fmt.Println("月:", d.Month())
	fmt.Println("日:", d.Day())
	fmt.Println("时:", d.Hour())
	fmt.Println("分:", d.Minute())
	fmt.Println("秒:", d.Second())
	fmt.Println("毫秒:", d.Nanosecond())

	//时间添加
	addT := d.Add(24 * time.Hour)
	fmt.Println(addT)
	//把指定的字符串 转成相应的对应日期格式类型 返回time类型
	parse, _ := time.Parse("2006-01-02 15:04:05", "2017-12-03 22:01:02")
	fmt.Println(parse)
	//时间格式化
	format := d.Format("2006-01-02 15:04:05.000 PM")
	fmt.Println(format)
}
