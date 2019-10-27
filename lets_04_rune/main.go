package main

import "fmt"

//go语言切片 字符串
func main() {
	s1 := "白萝卜"
	fmt.Println(s1)
	sArray := []rune(s1)        //把字符串通转换成 一个rune切片
	sArray[0] = '红'             //把rune 切片的第一个元素改成‘红’字 注意是char 字符类型
	fmt.Println(string(sArray)) //把rune切片转换成字符串
}
