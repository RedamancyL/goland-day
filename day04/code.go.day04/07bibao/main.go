package main

import "fmt"

// 闭包

func f1(f func()) {
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

// 定义一个函数对f2进行包装

// func f3(f func(int, int)) func() {
// 	tmp := func() {
// 		fmt.Println("你好")
// 		f()
// 	}
// 	return tmp
// }

func zzk(x func(int, int), m, n int) func() {
	tmp := func() {
		x(m, n)
	}
	return tmp
}

func main() {
	ret := zzk(f2, 100, 200)
	ret()
}
