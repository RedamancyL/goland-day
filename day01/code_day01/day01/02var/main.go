package main

import "fmt"

// 声明变量
// var name string
// var age int
// var isOk bool

// 批量声明
var (
	name string
	age  int
	isOk bool
)

func main() {
	name = "梦想"
	age = 18
	isOk = true
	// go语言中变量声明必须使用，不使用就编译不过去

	fmt.Print(isOk)             // 在终端中输出要打印的内容
	fmt.Println(age)            // 打印完指定内容之后会在后便加一个换行符
	fmt.Printf("name:%s", name) // 用于格式化输出

	// 声明变量同时赋值
	var s1 string = "lyw"
	fmt.Println(s1)

	// 类型推导（根据值判断该变量是什么类型）
	var s2 = "20"
	fmt.Println(s2)

	// 简短变量声明（只能在函数里使用）
	s3 := "哈哈哈"
	fmt.Println(s3)

	// s1 := "10" // 同一个作用域({})不能重复声明同名的变量
}
