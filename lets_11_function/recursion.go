package main

import "fmt"

// recursion of golang 递归
func main() {
	num := recursion(1)
	fmt.Println(num)
	fmt.Println("factorial=", factorial(2))
	fmt.Println("steps=", steps(4))
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
/*
该题目来源于牛客网《剑指offer》专题。
一只青蛙一次可以跳上1级台阶，也可以跳上2级。求该青蛙跳上一个n级的台阶总共有多少种跳法（先后次序不同算不同的结果）。
1阶：共1种跳法；
2阶：共2种跳法；
3阶：共3种跳法；
n阶：先跳1级，还剩n-1级，跳法是f(n-1)；先跳2级，还剩n-2级，跳法是f(n-2)，共f(n-1)+f(n-2)种跳法；
得出一个斐波那契函数。
Go语言实现：
方法一：递归
*/
func steps(n int) int {
	if n <= 0 {
		panic("没有台阶跳什么？？？")
	}
	if n == 1 { //如果就一条台阶肯定就一种跳法
		return 1
	}
	if n == 2 { //如果就两条台阶肯定就二种跳法 1.第一次可以跳一条台阶  2.第二次可以跳二条台阶 也可以反正过来 加起来反正都是2次
		return 2
	}

	return (n - 1) + (n - 2)
}
