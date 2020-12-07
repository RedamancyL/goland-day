package main

import "fmt"

// 常量
// 定义了常量之后不能修改
// 在程序运行期间不会改变的量
const pi = 3.1415926

// 批量声明常量
const (
	StatusOk = 200
	NotFound = 404
)

// 批量声明常量时，如果某一行声明后没有赋值，默认就和上一行一致

const (
	n1 = 100
	n2
	n3
)

// iota:类似枚举
const (
	a1 = iota //0
	a2        //1
	a3        //2
)

const (
	b1 = iota //0
	b2        //1
	_         //2
	b3        //3
)

const (
	c1 = iota //0
	c2 = 100  //100
	c3        //200
	c4 = iota //3
)

const (
	d1, d2 = iota + 1, iota + 2 // 一行 d1=0+1 , d2=0+2
	d3, d4 = iota + 1, iota + 2 // d3= 1+1 ,d4= 1+2
)

// 定义数量级
const (
	_  = iota             // 0给_直接仍了
	KB = 1 << (10 * iota) // 1左移十位变为 1000000000(二进制到十进制) == 1024
	MB = 1 << (10 * iota) // 1左移二十位变为 10^20(二进制到十进制) == 1024X1024
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	// pi = 111
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)

	fmt.Println(b1)
	fmt.Println(b2)
	//fmt.Println(_) // None
	fmt.Println(b3)

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(c4)

	fmt.Println(d1) // 1
	fmt.Println(d2) // 2
	fmt.Println(d3) // 2
	fmt.Println(d4) // 3
}
