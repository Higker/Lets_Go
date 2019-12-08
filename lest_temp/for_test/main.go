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
	for _, c := range "Hello世界!" {
		if len(string(c)) == 3 {
			count++
		}
	}
	fmt.Println("str里面中文字符的长度是:", count)
}
