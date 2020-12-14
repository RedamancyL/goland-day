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
	// 自己造一个切片也是生成一个数组，数组包装成一个切片
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s2), cap(s2))

	// 2. 由数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	s3 := a1[0:4] // 左闭右开
	fmt.Println(s3)

	// *切片的容量是指底层数组的容量
	fmt.Println(cap(s3))                                    // s3 的容量 7
	fmt.Printf("len(s3):%d cap(s3):%d\n", len(s3), cap(s3)) // len(s3):4 cap(s3):7

	s4 := a1[3:]           									// [7 9 11 13]                         
	// 底层数组从切片的第一个元素到最后的元素数量
	fmt.Println(cap(s4))                                    // s4 的容量 4
	fmt.Printf("len(s4):%d cap(s4):%d\n", len(s4), cap(s4)) // len(s4):4 cap(s4):4

	// 切片再切片
	s5 := s4[3:]
	fmt.Printf("len(s5):%d cap(s5):%d\n", len(s5), cap(s5)) // len(s5):1 cap(s5):1

	// 切片是引用类型，都指向了底层的一个数组。
	// s4
	fmt.Println("s4 ", s4)   // s4  [7 9 11 13]
	a1[6] = 1300 // 修改了底层数组的值 
	fmt.Println("s4 ", s4) 	 // s4  [7 9 11 1300]
}
