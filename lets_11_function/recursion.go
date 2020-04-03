package main

import "fmt"

// recursion of golang 递归
func main() {
	num := recursion(1)
	fmt.Println(num)
	fmt.Println("factorial=", factorial(2))
}

func recursion(n int) int {
	n++
	if n == 100 { //控制什么时候不递归
		return n
	}
	return recursion(n)

}

// factorial 阶乘 递归实现
func factorial(n int64) int64 {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1) // n = n - 1 * n -> n * n
}

// 2
// 2 = 2 -1
// 2 * 1
// 2
/*
	3! = 3*2*1 = 3*2!
	4! = 4*3*2*1 = 4*3!
	5! = 5*4*3*2*1 = 5*4!
	6! = 6*5*4*3*2*1 = 6*5!
*/
