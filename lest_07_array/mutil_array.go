package main

import "fmt"

/*
	使用二维数组输出下面内容
	0 0 0 1 0 0 0
	2 0 0 0 0 1 0
	0 0 3 0 0 1 0
	0 0 0 1 0 0 0
*/
func main() {
	var array [4][7]int
	array[0][3] = 1
	array[1][0] = 2
	array[1][5] = 1
	array[2][2] = 3
	array[2][5] = 3
	array[3][3] = 1
	//len获取1位数字长度
	for i := 0; i < len(array); i++ {
		//len(array[i]获取的二维数组长度)
		for j := 0; j < len(array[i]); j++ {
 			fmt.Print(array[i][j], " ")
		}
		fmt.Println("")
	}
}
