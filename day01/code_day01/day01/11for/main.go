package main

import "fmt"

func main() {
	// //基本格式
	// for i := 0; i < 10; i++ { // i++ 累计模式，在原来的基础上加1
	// 	fmt.Println(i)
	// }

	// // 变种1
	// var i = 5
	// for ; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// // 变种2
	// var i2 = 5
	// for i2 < 10 {
	// 	fmt.Println(i2)
	// 	i2++
	// }

	// // 无限循环
	// for {
	// 	fmt.Println("123")
	// }

	// for range循环
	s := "hello世界"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}

	// 流程控制之跳出for循环

	// 当i=5的时候跳出for循环
	fmt.Println("start")
	for i := 0; i < 10; i++ {
		if i == 5 {
			break // 跳出for循环
		}
		fmt.Println(i)
	}
	fmt.Println("over")

	// 当i=5时，跳过此次for循环(不执行for循环内部的打印语句)
	fmt.Println("start")
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue // 继续下一次循环
		}
		fmt.Println(i)
	}
	fmt.Println("over")
}
