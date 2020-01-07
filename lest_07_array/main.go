/*
 * @Author: Deencode
 * @Date: 2020-01-07 19:55:58
 * @LastEditors  : BinScholl
 * @LastEditTime : 2020-01-07 20:30:56
 * @Description: go语言中的数组使用  Array
 * @Github: https://github.com/Deencode
 */

package main

import "fmt"

//数组的使用

var (
	people [3]string
)

func main() {
	//给people赋值
	for i := 0; i < len(people); i++ {
		people[i] = fmt.Sprintf("小米%d", i)
	}
	for _, v := range people {
		fmt.Println(v)
	}
	//声明一个5位元素位int64类型的数组并且初始化默认值
	ages := [5]int64{11, 16, 18, 19, 22}
	for i := 0; i < len(ages); i++ {
		if i == 0 {
			fmt.Print("ages = [\t")
		}
		//遍历打印ages
		fmt.Print(ages[i], "\t")
		if i == len(ages)-1 {
			fmt.Println("]")
		}
	}
	//比较方便的声明方式  ’...‘代表你不确定这个数组有几个元素位 让go语言自己去推断
	isGood := [...]bool{true, false, true}
	fmt.Println(isGood)
	//没有初始化 数组会有默认值
	nums := [5]int{1, 2, 3}
	fmt.Println(nums)
	status := [3]bool{true}
	fmt.Println(status)

	//多维数组 并且 数组类型要一至
	n := [3][2]int{
		[2]int{1, 2},
		[2]int{2, 2},
		[2]int{3, 2},
	}
	fmt.Println(n)
	//访问n数组中的第二个元素里面的第二个元素
	fmt.Println(n[1][1]) //2
	fmt.Println()
	//通过下标去赋值 格式:下标:值
	floats := [10]float64{1: 1, 9: 9}
	fmt.Println(floats)
}
