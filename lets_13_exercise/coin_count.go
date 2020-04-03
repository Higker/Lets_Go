package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "adbdcdAKSWWOOA"
	begin := len(str) //先获取到原来的长度
	fmt.Println("str=", begin)
	tmp := strings.ReplaceAll(str, "d", "")
	fmt.Println(tmp)
	fmt.Printf("%T\n", tmp)
	num := len(tmp)
	fmt.Println(num)
	// // n := len(string(tmp)) //通过替换得到现在的长度
	en := (begin - num)     //然后把原来的长度减去现在的长度就知道这个字符出现了多少次了
	fmt.Println("len=", en) // 出现了3次
}
