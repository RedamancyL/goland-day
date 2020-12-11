package main

import "fmt"

// 数组

// 存放元素的容器
// 必须指定存放元素的类型和容量（长度）
// 数组长度是数组类型的一部分
func main() {
	var a1 [3]bool // [false false false]
	var a2 [4]bool //[false false false false]

	fmt.Printf("a1:%T a2:%T\n", a1, a2) // a1:[3]bool a2:[4]bool

	// 数组的初始化
	// 如果不初始化：默认元素都是零值（布尔值：false ，整型和浮点型都是0，字符串：“”）
	fmt.Println(a1, a2)

	// 1. 初始化方式1 (首先先需要定义)
	// var a3 [3]bool
	a3 := [3]bool{true, true, true}
	fmt.Println(a3)

	// 2. 初始化方式2
	// a10 := [9]int{0,1,2,3,4,5,6,7}
	// 根据初始值自动推断数组的长度是多少
	a10 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(a10)

	// 3. 初始化方式3
	a11 := [5]int{0: 1, 4: 2}
	fmt.Println(a11)

	// 数组的遍历
	city := [...]string{"北京", "上海", "深圳"}

	// 1.根据索引遍历
	for i := 0; i < len(city); i++ {
		fmt.Println(city[i])
	}

	// 2.for range遍历 (键值循环)
	for i, v := range city {
		fmt.Println(i, v)
	}

	//多纬数组
	//[[1 2][3 4][5 6]]
	var a23 [3][2]int
	a23 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a23)

	a24 := [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a24)

	// 多维数组的遍历
	for _, v := range a24 {
		// fmt.Println(i, v)
		for _, w := range v {
			fmt.Println(w)
		}
	}

	// 数组是值类型
	b1 := [3]int{1, 2, 3} // [1 2 3]
	b2 := b1              // [1 2 3 ](相当于python的深拷贝)
	b2[0] = 100           // b2:[100 2 3]
	fmt.Println(b1)       // b1:[1 2 3]
	fmt.Println(b2)       // b2:[100 2 3]

	// 练习
	// 1. 求数组[1, 3, 5, 7, 8]所有元素的和
	// 2. 找出数组中和为指定值的两个元素的下标，
	// 比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。

	arr1 := [...]int{1, 3, 5, 7, 8}
	num1 := 0
	for _, num := range arr1 {
		num1 += num
	}
	fmt.Println(num1)

	for i, x := range arr1 {
		for j, w := range arr1 {
			num2 := x + w
			if num2 == 8 {
				fmt.Printf("(%d %d)\n", i, j)
			} else {
				continue
			}

		}
	}

}
