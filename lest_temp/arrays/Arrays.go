package main

import "fmt"

var (
	arr1 = [5]int{}
	arr2 = [6]int{}
)

type city struct {
	name    string
	zipCode int
}

func main() {
	for i := range arr1 {
		arr1[i] = i
	}
	fmt.Println(arr1)
	//arr2 = arr1 类型不一致 因为go语言里面的数组的长度也是数组类型的一部分
	arr3 := arr1
	fmt.Println(arr2)
	fmt.Println(arr3 == arr1)

	var citys [3]city
	citys[0] = city{name: "上海", zipCode: 20000}
	citys[1] = city{"北京", 10001}
	citys[2] = city{"重庆", 20031}
	fmt.Println(citys)

	var total = [5]float64{1: 89.6, 4: 99.1}
	fmt.Println(total)

	var money = total
	money[0] = 100 //值拷贝 go的array是值类型
	fmt.Println(money)
	fmt.Println(total)
}
