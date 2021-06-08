## 变量类型
变量（Variable）的功能是存储数据。不同的变量保存的数据类型可能会不一样。 Go 拥有各值类型，包括字符串，整形，浮点型，布尔型等

## 变量声明
* Go语言中的变量需要声明后才能使用
* 同一作用域内不支持重复声明
* 并且Go语言的变量声明后必须使用

否则，编译时都会提示错误

### 可见性
```go
1）声明在函数内部，是函数的本地值，类似private
2）声明在函数外部，是对当前包可见(包内所有.go文件都可见)的全局值，类似protect
3）声明在函数外部且首字母大写是所有包可见的全局值,类似public
```


## 标准声明
Go语言的变量声明格式为：
```go
var 变量名 变量类型
```
变量声明以关键字var开头，变量类型放在变量的后面，行尾无需分号，如：
```go
var name string
var age int
var isOk bool
```
当然，如果需要批量申明，可以简化为：
```go
var(
    a1, a2 string
    b int
    c bool
    d float32
)
```

## 变量的初始化
Go语言在声明变量的时候，会自动对变量对应的内存区域进行初始化操作。每个变量会被初始化成其类型的默认值： 

* 整型和浮点型变量的默认值为0
* 字符串变量的默认值为空字符串
* 布尔型变量默认为false 
* 切片、函数、指针变量的默认为nil

变量初始化的标准格式：
```go
var 变量名 类型 = 表达式
```

如：
```go
var name string = "小明"
var age int = 12
var isOk bool = true
```

### 类型推导
Go语音支持将变量的类型省略，编译器会根据等号右边的值来推导变量的类型完成初始化，如：
```go
var name = "小明"
var age = 12
var isOk = true
```
我们也可以一次初始化多个变量：
```go
var name, age, isOk = "小明", 12, true
```

### 短变量声明
在函数内部，可以使用更简略的 := 方式声明并初始化变量:
```go
func foo() {
	name := "小明"
	age := 12
	isOk := true
    fmt.Println(name, age, isOk)
}
```

### 匿名变量
在使用多重赋值时，想要忽略某个值，可以使用匿名变量（`anonymous variable`）。 用下划线`_`表示，如：
```go
func foo() (int, string, bool {
    name := "小明"
    age := 12
    isOk := true
    return name, age, isOk
}
func main() {
    x, _, _ := foo()
    _, y, _ := foo()
    fmt.Println("name=", x)
    fmt.Println("age=", y)
}
```
匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明

> 注意：
```go
函数外的每个语句都必须以关键字开始（var、const、func等）
:= 不能使用在函数外
_ 多用于占位，表示忽略值
```
