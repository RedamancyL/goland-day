package main

import "fmt"

// defer
func deferDemo() {
	fmt.Println("start")

	// 释放资源的时候（等于python的__del__）
	defer fmt.Println("这是1") // 3 defer把它后面的语句延迟到函数即将返回的时候再执行
	defer fmt.Println("这是2") // 2  一个函数中可以有多个defer函数
	defer fmt.Println("这是3") // 1  后进先出

	fmt.Println("end")
}

func main() {
	deferDemo()
}
