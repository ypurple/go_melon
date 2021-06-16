## struct(结构体)
Go语言中没有“类”的概念，也不支持“类”的继承等面向对象的概念。Go语言中通过结构体的内嵌再配合接口比面向对象具有更高的扩展性和灵活性

结构体：
```go
1. 用来自定义复杂数据结构
2. struct里面可以包含多个字段（属性）
3. struct类型可以定义方法，注意和函数的区分
4. struct类型是值类型
5. struct类型可以嵌套
6. Go语言没有class类型，只有struct类型
7. 结构体是用户单独定义的类型，不能和其他类型进行强制转换
8. golang中的struct没有构造函数，一般可以使用工厂模式来解决这个问题。
9. struct中的每个字段可以增加tag。这个tag可以通过反射的机制获取到，最常用的场景就是json序列化和反序列化
```

### 结构体的定义
使用type和struct关键字来定义结构体:
```go
type 类型名 struct {
    字段名 字段类型
    字段名 字段类型
    …
}
```
如：
```go
type person struct {
    name string
    city string
    age  int
}
```

### 实例化
只有当结构体实例化时，才会真正地分配内存，才能使用结构体的字段。

```go
type person struct {
    name string
    city string
    age  int8
}

func main() {
    var p1 person
    p1.name = "pprof.cn"
    p1.city = "北京"
    p1.age = 18
    fmt.Printf("p1=%v\n", p1)  //p1={pprof.cn 北京 18}
    fmt.Printf("p1=%#v\n", p1) //p1=main.person{name:"pprof.cn", city:"北京", age:18}
}
```

- 键值对初始化：

```go
p5 := person{
    name: "pprof.cn",
    city: "北京",
    age:  18,
}
fmt.Printf("p5=%#v\n", p5) //p5=main.person{name:"pprof.cn", city:"北京", age:18}
```

- 值的列表初始化:
```go
p8 := &person{
    "pprof.cn",
    "北京",
    18,
}
fmt.Printf("p8=%#v\n", p8) //p8=&main.person{name:"pprof.cn", city:"北京", age:18}
```

### 匿名结构体

```go
package main

import (
    "fmt"
)

func main() {
    var user struct{Name string; Age int}
    user.Name = "pprof.cn"
    user.Age = 18
    fmt.Printf("%#v\n", user)
}
```

### 构造函数
Go语言的结构体没有构造函数，我们可以自己实现。 例如，下方的代码就实现了一个person的构造函数。 
结构体比较复杂的话，值拷贝性能开销会比较大，通常构造函数返回的是结构体指针类型。

```go
func newPerson(name, city string, age int8) *person {
    return &person{
        name: name,
        city: city,
        age:  age,
    }
}
```

调用构造函数：
```go
p9 := newPerson("pprof.cn", "测试", 90)
fmt.Printf("%#v\n", p9)
```

### 方法和接收者
Go语言中的方法（`Method`）是一种作用于特定类型变量的函数，也叫做接收者（`Receiver`）。
接收者的概念就类似于其他语言中的`this` 或者 `self`

```go
//Person 结构体
type Person struct {
    name string
    age  int8
}

//NewPerson 构造函数
func NewPerson(name string, age int8) *Person {
    return &Person{
        name: name,
        age:  age,
    }
}

//Dream Person做梦的方法
func (p Person) Dream() {
    fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

func main() {
    p1 := NewPerson("测试", 25)
    p1.Dream()
}
```

> 方法与函数的区别是，函数不属于任何类型，方法属于特定的类型。

