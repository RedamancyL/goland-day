package main

import "fmt"

// 函数

// 函数存在的意义？
// 函数是一段代码的封装
// 把一段逻辑抽象出来封装到一个函数中，
// 使用函数可以让代码更清晰，更加简洁

// 函数的定义

func sum(x int, y int) (ret int) {
	return x + y
}

// 没有返回值
func f1(x int, y int) {
	fmt.Println(x, y)
}

// 没有参数没有返回值
func f2() {
	fmt.Println("f2")
}

// 没有参数有返回值
func f3() (ret int) {
	return 3
}

// 命名返回值
// 参数可以命名也可以不命名
// 命名的返回值就相当于在函数中声明一个变量
func f4(x int, y int) (ret int) {
	ret = x + y
	return // 因为返回值已经命名，所有写ret和不写差不多
}

// Go语言中支持多个返回值
// 多个返回值
func f5() (int, string) {
	return 26, "你好"
}

// 参数的类型简写
func f6(x, y, z int, xx, yy string) (int, string) {
	return x, xx
}

// 可变长参数
// 可变长参数必须放在函数参数的最后
func f7(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y) // y是一个int类型的切片
}

// Go语言中函数没有默认参数这个概念

func main() {
	r := sum(1, 2)
	fmt.Println(r)
	f3 := f3()
	fmt.Println(f3)

	m, n := f5()
	fmt.Println(m, n)

	f7("起风了")
	f7("起风了", 1, 23, 4, 5, 6) // 调用函数

	// 在一个命名的函数中不能够再声明函数

}
