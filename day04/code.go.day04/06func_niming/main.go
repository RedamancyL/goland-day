package main

import "fmt"

// 匿名函数

func main() {

	// 匿名函数
	// 函数内部没有办法声明带名字的函数
	// 匿名函数一般用在函数内部调用函数时使用
	f1 := func(x, y int) {
		fmt.Println(x + y)
	}
	f1(10, 20)

	// 如果值是调用一次的函数，还可以简写成立立即执行函数
	func(x, y int) {
		fmt.Println("你好")
		fmt.Println(x + y)
	}(100, 200)
}
