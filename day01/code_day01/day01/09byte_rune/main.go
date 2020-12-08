package main

import "fmt"

func main() {
	s := "hello 世界"

	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()

	// len()求得时byte字节的数量
	n := len(s)

	fmt.Println(n)

	for _, c := range s {
		// fmt.Println(c)
		fmt.Printf("%c\n", c)
	}

	// 修改字符串
	s2 := "白萝卜"      //把字符强制转换成一个rune切片
	s3 := []rune(s2) // ===》 "白" "萝" "卜"
	s3[0] = '红'      // 双引号表示字符串 ，单引号标示字符
	fmt.Println(string(s3))

	c1 := "红" // string
	c2 := '红' // int32
	fmt.Printf("c1:%T c2:%T\n", c1, c2)

	c3 := "H"
	c4 := byte('H')
	fmt.Printf("c1:%T c2:%T\n", c3, c4)

	// 类型转换
	n1 := 10 //int
	var f float64
	f = float64(n1)
	fmt.Println(f)
	fmt.Printf("%T\n", f)
}
