package main

import "fmt"

// 函数类型

func f1() {
	fmt.Println("hello，zzk")
}
func f2() int {
	return 2
}

func f4(x, y int) int {
	return x + y
}

// 函数也可以作为参数的类型
func f3(x func() int) int {
	ret := x()
	return ret
}

// func f5(z func(int int) int) {
// 	ret := z
// 	return 5
// }

func ff(a, b int) int {
	return a + b
}

// 函数还可以作为返回值
// f5括号是他的参数 其中 x为参数的名字叫x ，返回值叫func() int ,后边是返回值的类型
func f5(x func() int) func(int, int) int {
	return ff
}

func main() {
	a := f1
	fmt.Printf("%T\n", a) // func() 没有参数，没有返回值的参数类型
	a2 := f2
	fmt.Printf("%T\n", a2) // func() int 没有参数，没有返回值的参数类型
	a3 := f3(f2)
	fmt.Println(a3)
	fmt.Printf("%T\n", a3) // unc(func() int) int int 没有参数，没有返回值的参数类型

	a4 := f4
	// fmt.Println(a4)
	fmt.Printf("%T\n", a4) // func(int, int) int 没有参数，没有返回值的参数类型
	f7 := f5(f2)
	fmt.Printf("%T\n",f7)
}
