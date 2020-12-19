package main

import "fmt"

func main() {
	var m1 map[string]int
	fmt.Println(m1 == nil) // 还没有初始化（没有再内存中开辟空间）

	m1 = make(map[string]int, 10) // 要估算好该map的容量，避免再程序运行期间再动态扩容

	m1["年龄"] = 18
	m1["身高"] = 180
	m1["体重"] = 140
	fmt.Println(m1)
	// m1["zzk"] = 18

	fmt.Println(m1["zzk"]) // 0 如果不存在这个key拿到对应值类型的零值

	// 约定成俗用ok接收bool值
	value, ok := m1["zzk"]
	if !ok {
		fmt.Println("没有此人")
	} else {
		fmt.Println(value) // 18
	}

	// map的遍历
	fmt.Println("----------------map的遍历-------------------")
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	// 只遍历 k
	for k := range m1 {
		fmt.Println(k)
	}
	// 只遍历 v
	for _, v := range m1 {
		fmt.Println(v)
	}

	// 删除  g
	delete(m1, "年龄")
	fmt.Println(m1)
}
