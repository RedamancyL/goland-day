package main

import "fmt"

func main() {
	// 切片的定义
	var s1 []int    // 定义一个存放int类型元素的切片
	var s2 []string // 定义一个存放string类型元素的切片
	fmt.Println(s1, s2)

	// 空 在go表示 你没有开辟内存空间
	fmt.Println(s1 == nil) // true
	fmt.Println(s2 == nil) // true

	//初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"陕西", "江苏", "浙江"}

	fmt.Println(s1 == nil) // false
	fmt.Println(s2 == nil) // false

	//长度和容量
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s2), cap(s2))

	// 2. 由数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	s3 := a1[0:4] // 左闭右开
	fmt.Println(s3)

	// *切片的容量是指底层数组的容量
	fmt.Println(cap(s3)) // s3 的容量 7 

}
