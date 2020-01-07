/*
 * @Author: Deencode
 * @Date: 2019-12-13 20:38:08
 * @LastEditors  : Deencode
 * @LastEditTime : 2020-01-06 13:36:39
 * @Description: go语言中的运算符 加减乘除法 取余
 * @Github: https://github.com/Deencode
 */

package main

import "fmt"

//go 语言运算符
func main() {
	/*
	 * 算数运算符:
	 * 加+ 减-  乘* 除/(除法取商)  取余% 自加加++ 自减减--
	 */
	add1 := 200
	add2 := 101
	fmt.Printf("%d + %d = %d\n", add1, add2, (add1 + add2))

	sub1 := 301
	sub2 := 200
	fmt.Printf("%d - %d = %d\n", sub1, sub2, (sub1 + sub2))

	multi1 := 100
	multi2 := 2
	fmt.Printf("%d * %d = %d\n", multi1, multi2, (multi1 * multi2))

	div1 := 100
	div2 := 2
	fmt.Printf("%d / %d = %d\n", div1, div2, (div1 / div2))

	div3 := 100
	div4 := 3
	//%代表取余     2个百分号 : ——>是因为%格式化缘故取余
	fmt.Printf("%d %% %d = %d\n", div3, div4, (div3 % div4))

	fmt.Println(add(0))
	fmt.Println(sub(1))

	//下面是关系运算符 '= > < !='
	a := false
	b := true
	c := 100
	d := 200
	fmt.Println(a == b) //在go语言中等于号只能比较2个值类型一样的
	//fmt.Println(b == c) invalid operation: b == c (mismatched types bool and int
	fmt.Printf("c的值类型是%T b的值类型是%T\n", c, b)
	fmt.Println(c >= d)
	fmt.Println(c <= d)
	fmt.Println(d > c)
	fmt.Println(c < d)
	fmt.Println(d != c)

	//下面逻辑运算符 '& &&  | ||  & -> AND | -> OR ! -> NOT'
	fraction := 86
	if fraction >= 90 && fraction <= 100 {
		fmt.Println("你的成绩真优秀")
	} else if (fraction >= 80 && fraction <= 90) || (fraction >= 70 && fraction <= 80) {
		fmt.Println("整体来说还行吧")
	} else {
		fmt.Println("渣渣你需要努力了~")
	}
	fmt.Println(!true) //取反

	//下面是位运算符
	/*
		 1. & : 按位与 (两位均为1才为1)
			5的二进制是 101
			2的二进制是  10
			第一轮 1和0   就一个1 那结果就是 0
			第二轮 0和1   也是一个1  那结果就是 0
			第三轮 1和一个空 (这里的空可以想象成0) 也是一个1 那结果就是 0
			则 000
	*/
	fmt.Println(5 & 2) //所以结果都是0

	/*
		2. | : 按位或 (两位均有一个为1则为1)
			5的二进制是 101
			2的二进制是  10
			第一轮 1和0   就一个1 那结果就是 1
			第二轮 0和1   也是一个1  那结果就是 1
			第三轮 1和一个空 (这里的空可以想象成0) 也是一个1 那结果就是 1
			则 111
	*/
	fmt.Println(5 | 2)                   //所以结果是 111(这3个1是二进制并非我们人类的十进制)
	result := 5 | 2                      //赋值一个变量
	fmt.Printf("result = %d \n", result) //通过go的format自动把这个3个1二进制转成十进制 结果7

	/*
		3. ^ : 按位异或 (两位不一样则为1)
			5的二进制是 101
			2的二进制是  10
			第一轮 1和0   1和0不一样 那结果就是 1
			第二轮 0和1    0和1不一样  那结果就是 1
			第三轮 1和一个空 1和0 (这里的空可以想象成0)  那结果就是 1
			则 111
	*/
	fmt.Printf("二进制 5 ^ 2 = %b \n", 5^2) //通过format手动转成二进制输出才能看到效果 111
	fmt.Println(5 ^ 2)                   //这里fmt输出的结果是十进制 7

	/*
		4. >> : 右移 (>>箭头后面代表是右移多少位)
	*/
	fmt.Println(5 >> 2) //5的二进制是101  右移2位 1|01 (|代表界限)然后去掉移除的2位 就是1
	/*
		5. << :左移  (<< 箭头后代表是左移多少位)
		左移 几位就是在101后面补零
	*/
	fmt.Println(5 << 2) //5的二进制是101 左移2位   10100| (|代表界限)然后去掉移除的2位 就是10 1的二进制就是20
}

/**
 * @description: 一个数加加
 * @param int64 number
 * @return: number add sum
 */
func add(num int64) (sum int64) {
	num++ //在go语言中只能单独语句  不能像Java那样放到等于号 右边 int num = 1++
	return num
}

//一个数减减
func sub(num int64) (result int64) {
	num--
	return num
}
