package main

import "fmt"

//logo banner
const (
	BANNER string = `
	┌─┐       ┌─┐ + +
    ┌──┘ ┴───────┘ ┴──┐++
    │                 │
    │       ───       │++ + + +
    ███████───███████ │+
    │                 │+
    │       ─┴─       │
    │                 │
    └───┐         ┌───┘
        │         │
        │         │   + +
        │         │
        │         └──────────────┐
        │                        │
        │                        ├─┐
        │                        ┌─┘
        │                        │
        └─┐  ┐  ┌───────┬──┐  ┌──┘  + + + +
          │ ─┤ ─┤       │ ─┤ ─┤
          └──┴──┘       └──┴──┘  + + + +
                 神兽保佑
                代码无BUG!
	`
)

func main() {
	fmt.Println(BANNER)
	//程序不自动退出
	var YouName string
	fmt.Println("请输入你的名字:")
	fmt.Scanln(&YouName)
	fmt.Println("你的名字是:", YouName)
	fmt.Scanln(&YouName)
}
