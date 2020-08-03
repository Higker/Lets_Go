package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	//启动一个文件服务器方便API测试
	go fileServe()
	// 获取一个默认路由 router
	r := gin.Default()
	// curl http://localhost:8080/
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"os":         runtime.GOOS,
			"go-version": runtime.Version(),
			"now_time":   time.Now().Format("2006-01-02 15:04:05"),
		})
	})
	// curl http://localhost:8080/user/ding/222
	r.GET("/user/:name/*action", func(c *gin.Context) {
		c.String(200, c.Param("name")+" / "+c.ClientIP())
	})
	// curl -i http://localhost:8080/user/ding\?id\=999
	r.GET("/user/:name", func(c *gin.Context) {
		c.String(200, c.Param("name")+" / ID="+c.DefaultQuery("id", "1"))
	})
	// http://localhost:8088/
	r.POST("/form", func(c *gin.Context) {
		name := c.DefaultPostForm("username", "DefaultName")
		ls := c.PostFormArray("language")
		fmt.Fprintln(c.Writer, name, ":", ls)
	})
	//fmt.Println(runtime.GOOS)
	r.Run(":8080")
}

func fileServe() {
	http.ListenAndServe(":8088", http.FileServer(http.Dir("./")))
}
