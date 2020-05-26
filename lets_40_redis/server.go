// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/5/18 - 1:04 下午

package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"html/template"
	"net/http"
	"strconv"
)

var (
	userList = make([]*user, 0, 10) //模拟数据库
	Rdb      *redis.Client
	zsetKey  = "ranking"
)

// 用户结构体
type user struct {
	Id       int
	Name     string
	Integral int
	PicUrl   string
}

// 初始化连接
func initClient() (err error) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "128.199.155.162:6379",
		Password: "deen.job", // no password set
		DB:       0,          // use default DB
	})
	_, err = Rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func init() {
	setData()
	err := initClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	//初始化一个zset列表
	var rankList []redis.Z
	//循环添加数据
	for _, user := range userList {
		rankList = append(rankList, redis.Z{Score: float64(user.Integral), Member: user.Name})
	}
	fmt.Println(rankList)
	fmt.Println("ADD ZSET")
	// 添加到redis db中
	Rdb.ZAdd(zsetKey, rankList...)
}
func main() {
	http.HandleFunc("/ranking_list", ranking)
	http.HandleFunc("/gift", gift)
	http.HandleFunc("/rank", rankList)
	http.ListenAndServe(":8080", nil)
}
func ranking(w http.ResponseWriter, req *http.Request) {
	files, err := template.ParseFiles("./ranking.tmpl")
	if err != nil {
		fmt.Sprintf("loading template file failed. %v\n", err)
	}
	files.Execute(w, userList)
}
func rankList(w http.ResponseWriter, req *http.Request) {
	result, err := Rdb.ZRevRangeWithScores(zsetKey, 0, 2).Result()
	if err != nil {
		fmt.Fprintf(w, func() string {
			marshal, _ := json.Marshal(err)
			return string(marshal)
		}())
		return
	}
	fmt.Fprintf(w, func() string {
		marshal, _ := json.Marshal(result)
		return string(marshal)
	}())

}
func gift(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	// 获取id
	id, _ := strconv.Atoi(req.URL.Query().Get("id"))
	if id > len(userList) {
		result := map[string]interface{}{"code": 500, "message": "你的id输入有误！！！"}
		fmt.Fprintf(w, func() string {
			marshal, _ := json.Marshal(result)
			return string(marshal)
		}())
		return
	}
	addGift(id)
	// 统一json返回结果
	result := map[string]interface{}{"code": 200, "integral": "添加100积分成功~"}
	fmt.Fprintf(w, func() string {
		marshal, _ := json.Marshal(result)
		return string(marshal)
	}())
}

// 设置静态数据
func setData() {
	userList = append(userList, &user{Id: 1, Name: "测试用户1", Integral: 0, PicUrl: `https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1589790107490&di=2d7bde662adb57ad326f6871f2efe170&imgtype=0&src=http%3A%2F%2F01.minipic.eastday.com%2F20170424%2F20170424110224_2976319ad50708e3ef72863260dd0536_4.jpeg`})
	userList = append(userList, &user{Id: 2, Name: "测试用户2", Integral: 0, PicUrl: `https://t1.hddhhn.com/uploads/tu/201810/415/055.jpg`})
	userList = append(userList, &user{Id: 3, Name: "测试用户3", Integral: 0, PicUrl: `https://t1.hddhhn.com/uploads/tu/201704/330/c1.jpg`})
}
func addGift(id int) {
	// 通过id那到名字
	var name string
	for _, u := range userList {
		if u.Id == id {
			name = u.Name
		}
	}
	// 然后把这个名字的用户的积分加100
	Rdb.ZIncrBy(zsetKey, 100, name)
}
