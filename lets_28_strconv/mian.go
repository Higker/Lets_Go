// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/4/21 - 2:51 下午

package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "10000"
	//通过strconv包的ParseInt方法就可以把一个字符串转成int类型，10代表十进制，32位int
	//这里32就是告诉转换器只使用了32位二进制位 使用返回是int64  然后自己再通过int32()强转即可得到int32类型
	//如果转入的0就是int类型
	num, _ := strconv.ParseInt(str, 10, 32)
	fmt.Printf("%v %T \n", num, num)
	num2, _ := strconv.Atoi(str) //array char 所有缩写成这样 也是字符串换成数字
	fmt.Printf("%v %T \n", num2, num2)
	num3 := 9999
	fmt.Printf("%v %T \n", strconv.Itoa(num3), strconv.Itoa(num3))
	Bstr := "true"
	fmt.Println(strconv.ParseBool(Bstr))
}
