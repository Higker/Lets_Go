package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//go语言中的map使用
func main() {
	//错误声明方法
	var m1 map[int]string
	fmt.Println("m1 == nil ?", m1 == nil) //m1 true 空的map 没有初始化内存地址
	//之前我们学过make内置函数 通过make函数创建
	m1 = make(map[int]string, 10)
	m1[0] = "Go"
	m1[1] = "C++"
	m1[2] = "Java"
	for k, v := range m1 {
		fmt.Printf("m1[%d] = %s \n", k, v)
	}

	_, notNull := m1[3] // 通过返回的第二参数判断map是否包含某元素

	if !notNull {
		m1[3] = "Python"
	}
	fmt.Println("m1[3] =", m1[3])

	//只遍历key
	for k := range m1 {
		fmt.Println(k)
	}
	//只遍历value
	for _, v := range m1 {
		fmt.Println(v)
	}
	//删除某个元素
	delete(m1, 1)
	fmt.Println(m1)

	students := make(map[string]int, 100)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		name := fmt.Sprintf("学生%d", i)
		fraction := rand.Intn(100)
		students[name] = fraction
	}
	fmt.Println(students)
	//对学生分数排序
	sortFraction(students)

	//手动造数据
	stu := map[string]int{"小名": 90, "小丁": 100, "小红": 60, "小王": 29}
	//fmt.Println(stu)
	sortFraction(stu)
}

//自己通过map加序号进行key-value颠倒排序
func sortFraction(students map[string]int) {
	olbStu := make(map[int]string, len(students)) //存储等下颠倒的学生数据
	//这里的操作是把分数换成key 方便下面进行通过分数排序
	for n, f := range students {
		olbStu[f] = n
	}
	frSli := make([]int, 0, len(olbStu))
	for f := range olbStu {
		frSli = append(frSli, f) //通过分数key 存储到 分数切片里面
	}
	sort.Ints(frSli)
	//用来存储 排序好的学生名字和分数 || 这种弊端就是怕元素位重复到 key重复赋值
	sortStu := make(map[string]int, len(students))
	for _, f := range frSli {
		//fmt.Println("f", f)
		stuName := olbStu[f] //通过分数去之前的那个 分数作为key的map里面找出学生名字
		sortStu[stuName] = f //然后在建一个学生名字的map 存储 格式为 姓名,分数 的map
		fmt.Println("sortStu[", stuName, "]", sortStu[stuName])
	}
	//fmt.Println(sortStu) //这里是因为map 的key 会默认排序 所以输出还是原来一样的
}
