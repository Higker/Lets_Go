package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//匿名结构体 只需要使用一次 并且里面的字段大写方便后面的json工具包进行调用
	var stuc = struct {
		Name string
		Age  int
	}{"NAME", 100}
	j, _ := json.Marshal(stuc) //将结构体序列号成json字符串
	fmt.Println(string(j))
}
