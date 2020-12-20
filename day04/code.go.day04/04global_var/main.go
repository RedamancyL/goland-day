package main

import "fmt"

// 变量的作用域

var x = 100 // 定义一个全局变量

// 定义一个函数
func f1() {
	// x := 10

	name := "你好"
	// 函数中查找变量的顺序
	// 1. 现在函数内部查找
	// 2. 找不到就往函数的外面查找，一直找到全局
	fmt.Println(x)
	fmt.Println(name)
}

func main() {
	f1()
	// fmt.Println(name) // 函数内部定义的变量只能在该函数内部使用

	// 语句块作用域
	if i := 100; i < 18 {
		fmt.Println("不要喝酒，不要抽烟")
	}
	// fmt.Println(i) // 访问一不到 i
	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}
	// fmt.Println(j) // 访问一不到 j

}
