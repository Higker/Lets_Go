package main

import "fmt"

//GO中的slice 扩容   append()方法
//https://www.liwenzhou.com/posts/Go/06_slice/
func main() {
	books := []string{"西游记", "水浒传", "红楼梦"}
	fmt.Println("books =", books, "len=", len(books), "cap=", cap(books))
	/*
		错误的写法
		index out of range
		这种写法会出现索引越界
		books[3] = "三国演义"
	*/
	//append函数返回值必须使用原来类型的变量来接受
	books = append(books, "三国演义")
	fmt.Println("books =", books, "len=", len(books), "cap=", cap(books))
	/*
		append扩容策略是:
		首先判断，如果新申请容量（cap）大于2倍的旧容量（old.cap），最终容量（newcap）就是新申请的容量（cap）。
		否则判断，如果旧切片的长度小于1024，则最终容量(newcap)就是旧容量(old.cap)的两倍，即（newcap=doublecap），
		否则判断，如果旧切片长度大于等于1024，则最终容量（newcap）从旧容量（old.cap）开始循环增加原来的1/4，即（newcap=old.cap,for {newcap += newcap/4}）直到最终容量（newcap）大于等于新申请的容量(cap)，即（newcap >= cap）
		如果最终容量（cap）计算值溢出，则最终容量（cap）就是新申请容量（cap）。
		需要注意的是，切片扩容还会根据切片中元素的类型不同而做不同的处理，比如int和string类型的处理方式就不一样。
	*/
	var citys = make([]string, 0, 5) //len 为0 就是说明这个切片目前里面没有默认值
	//var citys = make([]string, 3, 5) //len如果是3的话 切片前3个元素会有默认值
	fmt.Println("citys =", citys, "len=", len(citys), "cap=", cap(citys))
	citys = append(citys, "上海") // append 函数添加的元素只会在原切片数据的最后一个元素
	fmt.Println("citys =", citys, "len=", len(citys), "cap=", cap(citys))
	cityS := []string{"深圳", "北京", "天津", "重庆", "香港"}
	citys = append(citys, cityS...) //...代表拆分   把一个切片 类似for循环那种 一个一个取值添加5
	fmt.Println("citys = ", citys)
	fmt.Println("citys =", citys, "len=", len(citys), "cap=", cap(citys))
}

/* 想自己实现slice扩容的append方法 失败了
	old := [5]int{} //先声明一个默认的数组
	nums := []int
	nums = appendImpl(old,nums,6)
	fmt.Println(nums)
func appendImpl(old int[5],s []int, num int) []int{
	fmt.Println(old)
	s = old[:]
	if(cap(s) == len(old)){
		//二倍扩容
		old := [cap(s)*2]int
		//遍历原来不够用的旧切片
		i:=0
		for i;i<len(s);i++{
			old[i] = s[i]
		}
		//遍历完之后加入需要加的元素
		old[i+1] = num
		s = old[:]
	}else{
		s = append(s,num)
	}
	return s;
}
*/
