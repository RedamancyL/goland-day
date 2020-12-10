package main

import "fmt"

func main() {
	var (
		a = 5
		b = 2
	)

	// 算术运算符
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	a++ // 单独的语句，不能放在=的右边赋值 ==》 a=a+1
	b-- // 单独的语句，不能放在=的右边赋值 ==》 b=b-1
	fmt.Println(a)

	// 关系运算符
	fmt.Println(a == b) // Go语言时强类型，相同类型的变量才能比较
	fmt.Println(a != b) // 不等于 返回bool值
	fmt.Println(a >= b) // 大于等于
	fmt.Println(a > b)  //大于
	fmt.Println(a <= b) //小于等于
	fmt.Println(a < b)  // 小于

	//逻辑运算符
	// 如果年龄大于18岁并且小于60岁
	var age = 22
	if age > 18 && age < 60 {
		fmt.Println("加油上班")
	} else {
		fmt.Println("可以退休了")
	}

	// 如果年龄小于18岁或者大于60岁
	if age < 18 || age > 60 {
		fmt.Println("可以休息")
	} else {
		fmt.Println("需要上班")
	}

	// not 取反。原来为真就为假，原来为假就为真
	isword := false
	fmt.Println(isword)
	fmt.Println(!isword)

	// 位运算：针对的是二进制
	// 5的二进制：101 ==>0101
	// 2的二进制：010 ==>0010
	// 3的二进制：011 ==>0011

	// &：按位与 参与运算的两数各对应的二进位相与。（两位均为1才为1）
	fmt.Println(5 & 2) // 0000 =0
	fmt.Println(5 & 3) // 0001 =1

	// |:按位或  参与运算的两数各对应的二进位相或。（两位有一个为1就为1）
	fmt.Println(5 | 2) // 0111 =7
	fmt.Println(5 | 3) // 0111 =7

	// ^:按位异或  参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。（两位不一样则为1）
	fmt.Println(5 ^ 2) // 0111 =7
	fmt.Println(5 ^ 3) // 0110 =6

	// <<:将二进制位左移指定数  左移n位就是乘以2的n次方。“a<<b”是把a的各二进位全部左移b位，高位丢弃，低位补0。
	fmt.Println(5 << 2) // 0101 ==》00010100 = 20
	fmt.Println(5 << 3) // 0101 ==》00101000 = 32+8

	// >>将二进制位右移指定数  右移n位就是除以2的n次方。“a>>b”是把a的各二进位全部右移b位。
	fmt.Println(5 >> 1) // 0101 ==》0010 = 2
	fmt.Println(5 >> 2) // 0101 ==》0001 = 1

	// 左移太多超出范围 为0
	var m = int8(1)
	fmt.Println(m << 10) // 10000000

	var x int
	x = 1
}
