// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/5/22 - 5:58 下午

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/DING", func(writer http.ResponseWriter, request *http.Request) {
		files, err := template.ParseFiles("./form.tmpl")
		if err != nil {
			fmt.Println("页面加载失败")
			return
		}
		files.Execute(writer, nil)
	})
	http.HandleFunc("/post", post)
	fmt.Println("服务器已经启动.")
	http.ListenAndServe(":8080", nil)
}
func post(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.URL.Query().Get("username"))
	fmt.Println(request.URL.Query().Get("password"))
}
