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
	// http://localhost:8080/
	r.POST("/form", func(c *gin.Context) {
		name := c.DefaultPostForm("username", "DefaultName")
		ls := c.PostFormArray("language")
		fmt.Fprintln(c.Writer, name, ":", ls)
	})

	r.MaxMultipartMemory = 8 << 20

	r.POST("/upload", func(c *gin.Context) {

		forms, err := c.MultipartForm()
		fmt.Println(err)
		files := forms.File["fileList"]
		for _, file := range files {
			c.SaveUploadedFile(file, file.Filename)
		}
		fmt.Fprintln(c.Writer, "upload file ", len(files), " successful")
	})
	// curl http://localhost:8080/bindJson -X POST  -H "Content-Type: application/json" -d "{user:"YNN",password:"defaultPwd"}"
	r.POST("/bindJson", func(c *gin.Context) {
		var user User
		err := c.ShouldBindJSON(&user)
		fmt.Println(err)
		fmt.Fprintln(c.Writer, &user)
	})
	// Browser ： http://localhost:8080/Hello.html
	// HTML template
	r.LoadHTMLGlob("./template/*")
	r.GET("/Hello.html", func(c *gin.Context) {
		c.HTML(200, "Hello.tmpl", gin.H{
			"text": "Hello Gin Framework~",
		})
	})
	// Redirect uri
	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://github.com/higker")
	})
	// 中间件 处理请求时间计算
	r.Use(timer)
	// http://localhost:8080/timer
	r.GET("/timer", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		c.String(200, "处理完成")
	})
	// http://localhost:8080/cookie
	r.GET("/cookie", func(c *gin.Context) {
		value, err := c.Cookie("YNN")
		if err != nil {
			c.SetCookie("YNN", "DS live💖 YNN❤", 60, "/", c.Request.Host, false, true)
		}
		c.String(200, "cookie值是:"+value)
	})
	//fmt.Println(runtime.GOOS)
	r.Run(":8080")
}

type User struct {
	Name     string `json:"user" form:"name" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func fileServe() {
	http.ListenAndServe(":8088", http.FileServer(http.Dir("./")))
}
func timer(c *gin.Context) {
	start := time.Now()
	c.Next()
	total := time.Since(start)
	fmt.Println("处理请求耗费时间:", total)
}
