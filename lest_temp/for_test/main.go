package main

import "fmt"

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
}
