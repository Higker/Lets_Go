package main

import "fmt"

func main() {
	//"%v是获取值value缩写"
	str := "Hello"
	fmt.Printf("%v\n", str)
	var bin int = 1
	//%b是获取二进制值
	fmt.Printf("%b\n", bin)
	/*
		%T 查看类型
		%d 十进制
		%b 二进制
		%o 八进制
		%c 字符
		%s 字符串
		%v 获取值通用行
		%x 十六进制
		%f 浮点数
		%p 指针

		%+v 输出结构体带字段名
		%#v 值的go语言表示  例如字符串 打印值时会打印 “”号
		%% 百分号
	*/

	//scan函数获取用户输入
	var name string
	fmt.Scan(&name)
	fmt.Println("name = ", *(&name))

}
