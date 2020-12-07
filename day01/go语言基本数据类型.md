# 一. 基本数据类型
## 1.整型
整型分为以下两个大类： 按长度分为：int8、int16、int32、int64 对应的无符号整型：uint8、uint16、uint32、uint64

其中，uint8就是我们熟知的byte型，int16对应C语言中的short型，int64对应C语言中的long型。


| 类型        | 描述         |
| ------------- |:-------------:
|uint8	无符号   | 8位整型 (0 到 255)
|uint16 无符号   |16位整型 (0 到 65535)
|uint32	无符号   |32位整型 (0 到 4294967295)
|uint64	无符号   |64位整型 (0 到 18446744073709551615)
|int8	有符号   |8位整型 (-128 到 127)
|int16	有符号   |16位整型 (-32768 到 32767)
|int32	有符号   |32位整型 (-2147483648 到 2147483647)
|int64	有符号   |64位整型 (-9223372036854775808 到 9223372036854775807)


## 2.特殊整型

| 类型        | 描述         |
| ------------- |:-------------:
|uint	        |32位操作系统上就是uint32，64位操作系统上就是uint64
|int	        |32位操作系统上就是int32，64位操作系统上就是int64
|uintptr	    |无符号整型，用于存放一个指针


注意： 在使用int和 uint类型时，不能假定它是32位或64位的整型，而是考虑int和uint可能在不同平台上的差异。

## 3.浮点型

Go语言支持两种浮点型数：float32和float64

打印浮点数时，可以使用fmt包配合动词%f 如：
```
package main

import "fmt"

//浮点数

func main() {
	// math.MaxFloat32 // floot32最大值
	f1 := 1.23456
	fmt.Printf("%T\n", f1) // 默认Go语言中的小数都是float64类型
	f2 := float32(1.23456)
	fmt.Printf("%T\n", f2) // 显示声明float32类型
	// f1 = f2 类型不同不能相互赋值
}
```

## 4.复数 (一般不用)

complex64和complex128
```
var c1 complex64
c1 = 1 + 2i
var c2 complex128
c2 = 2 + 3i
fmt.Println(c1)
fmt.Println(c2)
```
复数有实部和虚部，complex64的实部和虚部为32位，complex128的实部和虚部为64位。

## 5. 布尔值

Go语言中以bool类型进行声明布尔型数据，布尔型数据只有true（真）和false（假）两个值。

```
package main

import "fmt"

func main() {
	b1 := true
	var b2 bool // 默认false
	fmt.Printf("%T\n", b1)
	fmt.Printf("%T value:%v\n", b2, b2)
}
```
注意：

1. 布尔类型变量的默认值为false。
2. Go 语言中不允许将整型强制转换为布尔型.
3. 布尔型无法参与数值运算，也无法与其他类型进行转换。