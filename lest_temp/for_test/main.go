package main

import (
	"fmt"
	"unicode"
)

//计算一个字符串里面包含有多少个中文字符
func main() {
	// var str string = "Hello世界!"
	// StrArray := []rune(str)
	// count := 0
	// for _, c := range StrArray {
	// 	if len(string(c)) == 3 {
	// 		count++
	// 	}
	// }
	count := 0
	for _, c := range "Hello World!您好世界~" {
		if len(string(c)) == 3 {
			count++
		}
	}
	fmt.Println("str里面中文字符的个数是:", count)

	//求2到100的素数

	for i := 2; i < 100; i++ {
		for j := 2; j < i; j++ {
			if i/j == 0 {
				fmt.Println(i)
				break
			}
		}
	}

	str := "ABCD上海Google公司"
	count2 := 0
	for _, v := range str {
		if len(string(v)) == 3 {
			count2++
		}
	}
	fmt.Println(count2)

	s := "这是个一包含汉字和英文的字符串, This is an apple"
	hzc := 0

	for _, v := range s {
		if unicode.Is(unicode.Han, v) {
			hzc++
		}
	}

	fmt.Printf("该字符串中找到中文个数是：%v\n", hzc)

}
