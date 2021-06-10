## map(集合)
Map是一种数据结构，用于存储一系列无序的键值对。它基于键存储的，键就像一个索引一样，可以快速快速检索数据，键指向与该键关联的值。

定义：
```go
var identifier map[keyType]valueType
```
> map的读取和设置也类似slice一样，通过key来操作；
> 
> 键必须是支持相等运算符 ("=="、"!=") 类型， 如 `number`、`string`、`pointer`、`array`、`struct`以及对应的 `interface`。
> 
> 值可以是任意类型，没有限制。

如：
```go
// 声明一个key是字符串，值为int的字典,这种方式的声明需要在使用之前使用make初始化
var numbers map[string]int
// 另一种map的声明方式
numbers := make(map[string]int)
numbers["one"] = 1  //赋值
numbers["ten"] = 10 //赋值
numbers["three"] = 3

fmt.Println("第三个数字是: ", numbers["three"]) // 读取数据
// 打印出来如:第三个数字是: 3
```

注意：
> map用散列表来实现，是无序的，每次打印出来的map都会不一样，它不能通过index获取，而必须通过key获取

> map的长度是不固定的，也就是和slice一样，也是一种引用类型

> 内置的len函数同样适用于map，返回map拥有的key的数量

> map的值可以很方便的修改，通过numbers["one"]=11可以很容易的把key为one的字典值改为11

> map和其他基本型别不同，它不是`thread-safe`，在多个`go-routine`存取时，必须使用`mutex lock`机制


### 初始化
#### 直接初始化
```go
var m1 map[string]float32 = map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
fmt.Printf("map m1 : %v\n", m1)
```

#### make初始化
```go
package main

import (
	"fmt"
)

func main() {
	// 创建了一个键类型为string,值类型为int的map
	m1 := make(map[string]int)
	// 也可以选择是否在创建时指定该map的初始存储能力，如创建了一个初始存储能力为5的map
	m2 := make(map[string]int, 5)

	m1["a"] = 1
	m2["b"] = 2
	fmt.Printf("局部变量 map m1 : %v\n", m1)
	fmt.Printf("局部变量 map m2 : %v\n", m2)
}
```

输出：
```go
局部变量 map m1 : map[a:1]
局部变量 map m2 : map[b:2]
```

### 基本操作
插入、更新、查找、删除、判断是否存在、求长度
```go
package main

import (
	"fmt"
)

func main() {
	m := map[string]string{"key0": "value0", "key1": "value1"}
	fmt.Printf("map m : %v\n", m)
	//map插入
	m["key2"] = "value2"
	fmt.Printf("inserted map m : %v\n", m)
	//map修改
	m["key0"] = "hello world!"
	fmt.Printf("updated map m : %v\n", m)
	//map查找
	val, ok := m["key0"]
	if ok {
		fmt.Printf("map's key0 is %v\n", val)
	}

	// 长度：获取键值对数量。
	len := len(m)
	fmt.Printf("map's len is %v\n", len)

	// cap 无效，error
	// cap := cap(m)    //invalid argument m (type map[string]string) for cap
	// fmt.Printf("map's cap is %v\n", cap)

	// 判断 key 是否存在。
	if val, ok = m["key"]; !ok {
		fmt.Println("map's key is not existence")
	}

	// 删除，如果 key 不存在，不会出错。
	if val, ok = m["key1"]; ok {
		delete(m, "key1")
		fmt.Printf("deleted key1 map m : %v\n", m)
	}
}
```

输出：
```go
map m : map[key0:value0 key1:value1]
inserted map m : map[key0:value0 key1:value1 key2:value2]
updated map m : map[key0:hello world! key1:value1 key2:value2]
map's key0 is hello world!
map's len is 3
map's key is not existence
deleted key1 map m : map[key0:hello world! key2:value2]
```

### map遍历
不能保证迭代返回次序，通常是随机结果，具体和版本实现有关。

```go
package main

import (
	"fmt"
)

func main() {
    scoreMap := make(map[string]int)
    scoreMap["张三"] = 90
    scoreMap["小明"] = 100
    scoreMap["王五"] = 60
    for k, v := range scoreMap {
        fmt.Println(k, v)
    }
}
```

map排序：先获取所有key，把key进行排序，再按照排序好的key，进行遍历。
```go
package main

import (
	"fmt"
	"sort"
)

func main() {
    dict := map[string]int{"张三": 90, "小明": 100, "王五": 60}
    var names []string
    for name := range dict {
        names = append(names, name)
    }
    sort.Strings(names) //排序
    for _, key := range names {
        fmt.Println(key, dict[key])
    }
}
```

### 函数间传递Map
函数间传递Map是不会拷贝一个该Map的副本的，也就是说如果一个Map传递给一个函数，该函数对这个Map做了修改，那么这个Map的所有引用，都会感知到这个修改。
```go
package main

import (
	"fmt"
)

func main() {
    dict := map[string]int{"张三": 90, "小明": 100, "王五": 60}
    modify(dict)
    fmt.Println(dict["张三"])
}
func modify(dict map[string]int) {
    dict["张三"] = 10
}
```

输出：
```go
10
```