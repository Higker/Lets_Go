/*
 * @Author: BinScholl
 * @Date: 2020-01-03 19:58:26
 * @LastEditors  : BinScholl
 * @LastEditTime : 2020-01-03 20:20:56
 * @Description: go语言中goto关键字的使用
 * @Github: https://github.com/BinScholl
 */

package main

import "fmt"

func main() {
	gotoDemo2()
	var breakFlag bool
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 9 {
				// 设置退出标签
				breakFlag = true
				break
			}
			fmt.Printf("%v-%v\n", i, j)
		}
		// 外层for循环判断
		if breakFlag {
			break
		}
	}
}
func gotoDemo2() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 9 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
	// 标签
breakTag:
	fmt.Println("结束for循环")
}
