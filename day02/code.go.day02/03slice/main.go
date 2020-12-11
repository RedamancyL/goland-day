package main

import "fmt"

func main() {
	// 切片的定义
	var s1 []int    // 定义一个存放int类型元素的切片
	var s2 []string // 定义一个存放string类型元素的切片
	fmt.Println(s1, s2)

	// 空 在go表示 你没有开辟内存空间
	fmt.Println(s1 == nil) 
	fmt.Println(s2 == nil)

	//初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"陕西", "江苏", "浙江"}

	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
}
