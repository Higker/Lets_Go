package main

import "fmt"

//slice 练习题
func main() {
	var s1 = make([]int, 3, 5)
	for i := 0; i < 5; i++ {
		s1 = append(s1, i)
	}
	fmt.Println("s1 = ", s1) //0001234

	ss := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("ss=", ss)
	fmt.Println("ss len=", len(ss), "cap", cap(ss))
	ss = append(ss[:3], ss[4:]...)
	fmt.Println("ss=", ss)
	//cap=8 是应为切片存储数据是底层数组存储的 所以移除了一个元素 只是底层数组的对应值被后面一个元素覆盖了 len是不变的
	fmt.Println("ss len=", len(ss), "cap", cap(ss))
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println("arr = ", arr)
	sss := append(arr[:3], arr[4:]...)
	fmt.Println("sss=", sss)
	fmt.Println("arr len=", len(arr), "arr=", arr)
	fmt.Println("sss len=", len(sss), "cap", cap(sss))
}

/*
	[Running] go run "d:\Go_workspace\src\Lets_Go\lets_08_slice\slice_practice.go"
	s1 =  [0 0 0 0 1 2 3 4]
	ss= [1 2 3 4 5 6 7 8]
	ss len= 8 cap 8
	ss= [1 2 3 5 6 7 8]
	ss len= 7 cap 8
	[Done] exited with code=0 in 2.487 seconds
*/
