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
}
