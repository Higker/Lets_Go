/*
 * @Author: Deencode
 * @Date: 2020-01-03 18:27:04
 * @LastEditors  : Deencode
 * @LastEditTime : 2020-01-03 19:46:05
 * @Description: go语言中的switch语句使用示例
 * @Github: https://github.com/Deencode
 */

package main

import "fmt"

func main() {
	switchDemo1(22)
	switchDemo2(11)
	switchDemon3(5)
	switchDemon4()
}

func switchDemo1(finger int) {
	switch finger {
	case 1:
		fmt.Println(finger, "是大拇指")
	case 2:
		fmt.Println(finger, "是食指")
	case 3:
		fmt.Println(finger, "是中指")
	case 4:
		fmt.Println(finger, "是无名指")
	case 5:
		fmt.Println(finger, "是小拇指")
	default:
		fmt.Println("你是个外星人吧")
	}
}
func switchDemo2(num int) {
	switch num {
	case 1, 3, 5, 7, 9:
		fmt.Println(num, "是奇数")
	case 2, 4, 6, 8:
		fmt.Println(num, "是偶数")
	default:
		//👆上面这种写法 要写到死吧
		// 通过下面这种方式计算 方便多了
		if num%2 == 0 {
			fmt.Println(num, "是偶数")
		}
		fmt.Println(num, "是奇数")
	}

}
func switchDemon3(age int) {
	switch {
	case age >= 0 && age <= 6:
		fmt.Println(age, "岁小屁孩一个")
	case age >= 7 && age <= 23:
		fmt.Println(age, "岁好好上学吧,不上学你会后悔的")
	case age >= 24 && age <= 40:
		fmt.Println(age, "岁成家立业了")
	case age > 40:
		fmt.Println(age, "岁你可以富可敌国了,千亿富豪了")
	}
}
func switchDemon4() {
	//fallthrough语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的。
	switch num := 75; { //num is not a constant
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200", num)
	}
}
