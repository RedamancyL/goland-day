## Go语言基础之map
Go语言中提供的映射关系容器为map，其内部使用散列表（hash）实现。

# map
map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用。

## map定义
Go语言中 map的定义语法如下：


```
map[KeyType]ValueType

其中，
KeyType:表示键的类型。
ValueType:表示键对应的值的类型。

```
map类型的变量默认初始值为nil，需要使用make()函数来分配内存。语法为：

```make(map[KeyType]ValueType, [cap])```
```
func main() {
	var m1 map[string]int
	fmt.Println(m1 == nil)   // 还没有初始化（没有再内存中开辟空间）

	m1 = make(map[string]int, 10)  // 要估算好该map的容量，避免再程序运行期间再动态扩容 

	m1["年龄"] = 18
	m1["身高"] = 180
	fmt.Println(m1)
}

```

其中cap表示map的容量，该参数虽然不是必须的，但是我们应该在初始化map的时候就为其指定一个合适的容量。

## map基本使用

map中的数据都是成对出现的，map的基本使用示例代码如下：
```
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
}
输出：

rue
map[体重:140 年龄:18 身高:180]
0
没有此人
```
map也支持在声明的时候填充元素，例如：

func main() {
	userInfo := map[string]string{
		"username": "沙河小王子",
		"password": "123456",
	}
	fmt.Println(userInfo) //
}

## 判断某个键是否存在
Go语言中有个判断map中键是否存在的特殊写法，格式如下:

```value, ok := map[key]```
举个例子：

```
func main() {
	var m1 map[string]int
	fmt.Println(m1 == nil) // 还没有初始化（没有再内存中开辟空间）

	m1 = make(map[string]int, 10) // 要估算好该map的容量，避免再程序运行期间再动态扩容

	m1["年龄"] = 18
	m1["身高"] = 180
	m1["体重"] = 140(map[string][]i
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
}
```
## map的遍历
Go语言中使用for range遍历map。

但我们只想遍历key的时候，可以按下面的写法：

```
func main() {
	var m1 map[string]int
	fmt.Println(m1 == nil) // 还没有初始化（没有再内存中开辟空间）

	m1 = make(map[string]int, 10) // 要估算好该map的容量，避免再程序运行期间再动态扩容

	m1["年龄"] = 18
	m1["身高"] = 180
	m1["体重"] = 140
	fmt.Println(m1)

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

    delete(m1, "年龄")
	fmt.Println(m1)
}
```
注意： 遍历map时的元素顺序与添加键值对的顺序无关。

## 使用delete()函数删除键值对
使用delete()内建函数从map中删除一组键值对，delete()函数的格式如下：

```delete(map, key)```
其中，

- map:表示要删除键值对的map
- key:表示要删除的键值对的键
```
	delete(m1, "年龄")
	fmt.Println(m1)
```

## 按照指定顺序遍历map
```
func main() {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("student %02d", i) //生成stu开头的字符串
		value := rand.Intn(100)             //生成0~99的随机整数
		scoreMap[key] = value
	}

	fmt.Println(scoreMap)
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}
```

## 元素为map类型的切片
下面的代码演示了切片中的元素为map类型时的操作：
```
func main() {
	// 元素类型为map的切片
	var s1 = make([]map[int]string, 3)
	fmt.Println(len(s1))              // 3
	// 没有对内部的map做初始化
	s1[0] = make(map[int]string, 2)  // 在s1[0]中再次初始化切片
	fmt.Println(s1[0])               // map[]
	fmt.Println(len(s1[0]))          // 0
	s1[0][10] = "程序001"
	s1[0][11] = "程序002"
	s1[0][12] = "程序003"
	fmt.Println(s1[0])              // map[10:程序001 11:程序002 12:程序003]

	fmt.Println(s1)                 // [map[10:程序001 11:程序002 12:程序003] map[] map[]]

输出：
3
map[]
0
map[10:程序001 11:程序002 12:程序003]
[map[10:程序001 11:程序002 12:程序003] map[] map[]]

```

## 值为切片类型的map

```
func main() {
	var m1 = make(map[string][]int, 10)
	m1["北京"] = []int{10, 20, 30}
	fmt.Println(m1)
}
```