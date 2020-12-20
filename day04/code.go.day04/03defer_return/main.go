package main

import "fmt"

// Go语言中函数的return不是原子操作，在底层是分为两步来执行
// 第一步：返回值赋值
// 第二步：真正ret返回
// 函数中如果存在defer，那么defer执行的时是在第一步和第二步之间

func f1() int {
	x := 5
	defer func() {
		x++ // 修改的x 不是返回值int
	}()

	return x
}

func f2() (x int) { // 先赋予x返回值5
	defer func() {
		x++ // 修改的x  加1
	}()

	return 5 // 返回值=x
}

func f3() (y int) { //
	x := 5
	defer func() {
		x++ //  修改的是x
	}()

	return x // 返回值 = y = x = 5
}

func f4() (x int) { //
	defer func(x int) {
		x++ // 改变的是函数中的副本
	}(x)

	return 5 // 返回值 = x = 5
}

func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())

}
