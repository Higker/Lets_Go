package main

import (
	"fmt"
)

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

//1.x =1 y=2
//2.会查看defer里面的calc如果参数是函数的话就执行参数对应的那个函数
//3. "A" 1 2 3
//4. 经过上面的执行 cals("AA",1,3)
//5. x = 10
//6. "B" 10 2 12
//7. cals("BB",10,12)
//8. y = 20 用来混淆人的思路的不用管就是给普通赋值
//9. "BB" 10 12 22
//10. "AA" 1 3 4
/*
	预测结果是:
	"A" 1 2 3
	"B" 10 2 12
	"BB" 10 12 22
	"AA" 1 3 4
*/
func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}
