package main

import (
	"fmt"
	"strings"
	"unicode"
)

//练习题 判断一个字符串里面有多少个中文
func main() {
	var tempStr string = "今天我学习了Go语言了."
	count := 0
	for _, v := range tempStr {
		if unicode.Is(unicode.Han, v) {
			count++
		}
	}
	fmt.Println("count = ", count)

	//计算字符串中的每个单词出现的次数
	wordStr := "How do you do"              //声明一个英文字符串遍历
	sliceStr := strings.Split(wordStr, " ") //通过split函数切割字符串
	cmap := make(map[string]int, 20)        //声明一个map存储 记录英文单词出现的次数
	for _, v := range sliceStr {
		if _, not := cmap[v]; !not {
			cmap[v] = 1
		} else {
			cmap[v]++
		}
	}
	fmt.Println(cmap)

	//回文判断
	//上海自来水来自海上
	//黄山落叶松叶落山黄

	tstr := "上海自来水来自海上"
	rstr := []rune(tstr)
	for i := 0; i < len(rstr); i++ {
		//比较方式就是 通过index 和最后一个end_index进行比较 首位比较
		fmt.Println(string(rstr[i]))
		for j := len(rstr); j >= i; j-- {
			if string(rstr[i]) == string(rstr[(j-1)]) {
				fmt.Printf("rstr[%d] = %s | rstr[%d] \n", i, string(rstr[i]), j, string(rstr[j]))
			}
		}
	}
	//通过上面案例 实现类似于 QQ敏感词汇拦截
	// str := "zyh 我操你妈的 ~"
	// mgc := "草泥马操你妈傻逼"
	// for _, v := range []rune(str) {
	// 	if unicode.In(v, mgc) {
	// 		fmt.Println("敏感词")
	// 	}
	// }
}
