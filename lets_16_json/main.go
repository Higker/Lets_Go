package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//go语言json数据转换 & 结构体字段必须能让外部包访问到
type person struct {
	Name string
	Age  int
}
type student struct {
	person
	Class    string
	Learning string
}

func main() {
	HttpServer()

}

func HttpServer() {
	http.HandleFunc("/user", getUser)
	http.ListenAndServe(":8080", nil)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	s1 := student{
		person{
			"宁宁",
			18,
		},
		"5班",
		"Golang",
	}
	fmt.Println("s1 = ", s1)
	json, err := json.Marshal(s1)
	if err != nil {
		fmt.Println(json)
	}
	fmt.Fprintln(w, string(json))
}
