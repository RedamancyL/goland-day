package main

import "fmt"

// copy

func main() {
	a1 := []int{1, 3, 5}
	a2 := a1 // 赋值
	var a3 = make([]int, 3, 5)
	copy(a3, a1)            // copy
	fmt.Println(a1, a2, a3) //[1 3 5] [1 3 5] [1 3 5]
	a1[0] = 100
	fmt.Println(a1, a2, a3) //[100 3 5] [100 3 5] [1 3 5]

	// 将a1中的索引为1的3这个元素删掉
	a1 = append(a1[:1], a1[2:]...)  // 添加另一个切片中的元素（后面加…）
	fmt.Println(a1)      // [100 5]
	fmt.Println(cap(a1)) // 3

	x1 := [...]int{1, 3, 5, 7, 9}
	s1 := x1[:]
	fmt.Println(s1, len(s1), cap(s1)) // [1 3 5 7 9] 5 5

	// 1. 切片不保存具体的值
	// 2. 切片对应一个底层数组
	// 3. 底层数组都是占用一块连续的内存
	s1 = append(s1[:1], s1[3:]...)    // 修改的底层数组
	fmt.Println(s1, len(s1), cap(s1)) // [1 7 9] 3 5

	// 通过切片可以修改底层数组
	s1[0] = 100     // 修改底层数组
	fmt.Println(x1) //[1 7 9 7 9]
}
