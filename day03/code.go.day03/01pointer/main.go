package main

import "fmt"

func main() {
	// 1. &: 取地址
	// 2. *: 根据地址取值
	n := 18
	p := &n
	fmt.Println(&n)       // 0xc0000b6010
	fmt.Printf("%T\n", p) // *int : int类型的指针
	// 2. *:根据地址取值
	m := *p
	fmt.Println(m)        // 18
	fmt.Printf("%T\n", m) // int

	var a1 *int //nil pointer

	fmt.Println(a1)
	var a2 = new(int) // new函数申请一个内存地址
	fmt.Println(a2)

	// var a = new(int)
	// *a = 100
	// fmt.Println(*a)

	// new函数申请一个内存地址
}
