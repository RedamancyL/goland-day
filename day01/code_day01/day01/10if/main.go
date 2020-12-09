package main

import "fmt"

// if 条件判断
func main() {
	// age := 36
	// if age > 18 {
	// 	fmt.Println("可以去酒吧了")
	// } else {
	// 	fmt.Println("你再等等！")
	// }

	// // 多个判断条件
	// if age > 35 {
	// 	fmt.Println("年纪大了")
	// } else {
	// 	fmt.Println("你还年轻")
	// }

	// 作用域
	// age变量此时只在if条件判断语句中生效
	// 执行完就销毁，减少内存占用
	if age := 17; age > 18 {
		fmt.Println("可以去酒吧了")
	} else {
		fmt.Println("你再等等！")
	}

	// fmt.Println(age)  //在这里时找不到age

}
