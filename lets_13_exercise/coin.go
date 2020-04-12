package main

import (
	"fmt"
	"strings"
)

/*
!!!本例子是我自己改编的!!!
//原例子是每包含一个就要处理:https://www.liwenzhou.com/posts/Go/09_function/
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中包含1个'e'或'E'分1枚金币
b. 名字中包含1个'i'或'I'分2枚金币
c. 名字中包含1个'o'或'O'分3枚金币
d: 名字中包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
	table        = map[string]int{"E": 1, "I": 2, "O": 3, "U": 4}
)

func main() {
	addUsers()
	left := dispatchCoin()
	fmt.Println("剩下：", left)
	fmt.Println(distribution)
}

func dispatchCoin() int {
	for name := range distribution { //先遍历拿出来每个用户map
		UpName := strings.ToUpper(name) // 把用户的名字全部大写 因为这里条件小和大小都一样
		for Key := range table {        //遍历
			if strings.Contains(UpName, Key) { //如果名字里面包含了某个字符
				switch Key { //对应的key做对应的事情
				case "E":
					distribution[name] += table[Key] //从table表里面找出对应的数量关系然后加起来
					//coins -= distribution[name]      //每一轮完成之后就要从总coins里面减去已经发的数量
				case "I":
					distribution[name] += table[Key]
					//coins -= distribution[name]
				case "O":
					distribution[name] += table[Key]
					//coins -= distribution[name]
				case "U":
					distribution[name] += table[Key]
					//coins -= distribution[name]
				}
			}
		}
	}
	return sumLeft()
}

func addUsers() {
	//把切片里面元素加到map里面
	for _, name := range users {
		distribution[name] = 0
	}
}
func sumLeft() (Left int) {

	for n := range distribution {
		Left += distribution[n]
	}
	Left = coins - Left
	return Left
}
