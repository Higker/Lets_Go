package main

import (
	"fmt"
)

//go语言中指针的使用
func main() {
	// 我们都知道，变量是一种使用方便的占位符，用于引用计算机内存地址。
	// Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。
	// 以下实例演示了变量在内存中地址：
	i := 100
	fmt.Println(&i) //go语言中取变量在内存地址使用 ‘&’ 符合
	// 什么是指针
	// 一个指针变量指向了一个值的内存地址。
	// 类似于变量和常量，在使用指针前你需要声明指针。指针声明格式如下：
	// 指针声明格式  var 指针名称 *指针类型
	var iptf *int     // 整型指针
	var fptf *float64 // 浮点型指针
	fmt.Println(iptf)
	fmt.Println(fptf)
	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */
	ip = &a        /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)
	fmt.Println("a == ip ？", &a == ip)
	n := 200
	ip = &n
	//在指针类型前面加上 * 号（前缀）来获取指针所指向的内容
	fmt.Println(ip)
	fmt.Println("a == ip ？", &a == ip)
	// 	Go 空指针
	// 当一个指针被定义后没有分配到任何变量时，它的值为 nil。

	// nil 指针也称为空指针。

	// nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。

	// 一个指针变量通常缩写为 ptr。

	// 查看以下实例：
	var bptr *bool
	var sptr *string
	var fptr *float32
	var iptr *int
	fmt.Printf("bptr 的值为 : %x\n", bptr)
	fmt.Printf("sptr 的值为 : %x\n", sptr)
	fmt.Printf("fptr 的值为 : %x\n", fptr)
	fmt.Printf("iptr 的值为 : %x\n", iptr)
	//空指针判断：
	fmt.Println("sptr != nil ?", sptr != nil)
	fmt.Println("iptr == nil ?", iptr == nil)
}

// [Running] go run "d:\Go_workspace\src\Lets_Go\lest_09_pointer\pointer.go"
// 0xc000066090
// <nil>
// <nil>
// a 变量的地址是: c000066098
// ip 变量储存的指针地址: c000066098
// *ip 变量的值: 20
// a == ip ？ true
// 0xc0000660c8
// a == ip ？ false
// bptr 的值为 : 0
// sptr 的值为 : 0
// fptr 的值为 : 0
// iptr 的值为 : 0
// sptr != nil ? false
// iptr == nil ? true

// [Done] exited with code=0 in 6.739 seconds
