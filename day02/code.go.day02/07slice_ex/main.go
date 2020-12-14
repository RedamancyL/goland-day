package main

import (
	"fmt"
	"sort"
)

func main() {
	var a = make([]string, 5, 10) // 创建切片，长度为5，容量为10
	fmt.Println(a, cap(a))
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a, cap(a))

	var a1 = [...]int{3, 7, 8, 9, 1}
	sort.Ints(a1[:]) // 对切片进行排序
	fmt.Println(a1)  // [1 3 7 8 9]
}
