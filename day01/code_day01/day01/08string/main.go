package main

import (
	"fmt"
	"strings"
)

// 字符串

func main() {
	// \ 本来就是具有特殊含义，我应该告诉程序我些的\就是一个单纯的\
	path := "/home/admin/passwd"

	fmt.Printf(path)

	s := "I'm ok"
	fmt.Printf(s)

	s2 := `
	  好好学习
	  天天向上
	`

	fmt.Printf(s2)

	s3 := "/home/admin/passwd"

	fmt.Printf(s3)

	// 字符串相关操作
	fmt.Println(len(s3))

	// 字符串拼接
	name := "理想"
	world := "lsp"

	ss := name + world
	fmt.Printf(name + world)
	fmt.Printf(ss)

	ss1 := fmt.Sprintf("%s%s", name, world) //Sprintf 拼接后返回
	fmt.Printf("%s%s", name, world)         // 直接打印到终端
	fmt.Println(ss1)

	// 分割
	ret := strings.Split(s3, "/")
	fmt.Println(ret)

	// 包含
	fmt.Println(strings.Contains(ss, "理性"))
	fmt.Println(strings.Contains(ss, "lsp"))

	// 前缀
	fmt.Println(strings.HasPrefix(ss, "lsp"))
	//后缀
	fmt.Println(strings.HasSuffix(ss, "lsp"))

	ss4 := "abcdec"
	fmt.Println(strings.Index(ss4, "c"))     //第一次出现的位置
	fmt.Println(strings.LastIndex(ss4, "c")) //最后一次出现的位置

	// 拼接
	fmt.Println(strings.Join(ret, "+"))
}
