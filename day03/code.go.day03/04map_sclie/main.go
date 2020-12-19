package main

import "fmt"

func main() {
	// // 元素类型为map的切片
	// var s1 = make([]map[int]string, 3)
	// fmt.Println(len(s1))
	// // 没有对内部的map做初始化
	// s1[0] = make(map[int]string, 2)
	// fmt.Println(s1[0])
	// fmt.Println(len(s1[0]))
	// s1[0][10] = "程序001"
	// s1[0][11] = "程序002"
	// s1[0][12] = "程序003"
	// fmt.Println(s1[0])

	// fmt.Println(s1)

	// 值为切片类型的map
	var m1 = make(map[string][]int, 10)
	m1["北京"] = []int{10, 20, 30}
	fmt.Println(m1)

	// var mapSlice = make([]map[string]string, 3)
	// for index, value := range mapSlice {
	// 	fmt.Printf("index:%d value:%v\n", index, value)
	// }
	// fmt.Println("after init")
	// // 对切片中的map元素进行初始化
	// mapSlice[0] = make(map[string]string, 10)
	// mapSlice[0]["name"] = "小王子"
	// mapSlice[0]["password"] = "123456"
	// mapSlice[0]["address"] = "沙河"
	// for index, value := range mapSlice {
	// 	fmt.Printf("index:%d value:%v\n", index, value)
	// }
}
