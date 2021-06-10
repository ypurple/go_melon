## 常量
常量，是在程序编译阶段就确定下来的值，程序在运行时无法改变该值。 Go 支持字符、字符串、布尔和数值 常量。

### 常量声明
常量的声明和变量声明非常类似，只是把`var`换成了`const`
```go
const 变量名 变量类型 = 值
```

#### 常规声明：
同样的，编译器会根据赋值来推导常量的类型，通常常量使用大写+下划线声明：
```go
const COLOR_RED = 0
const COLOR_GREEN = 1
const COLOR_BLUE = 2
```

#### 一起声明：
```
const(
    COLOR_RED = 0
    COLOR_GREEN = 1
    COLOR_BLUE = 2
)
```

> 常量在定义的时候**必须赋值**，如不提供类型和初始化值，那么视作与上一个常量相同

```go
const (
	U_STR_S = "abc"
    U_STR_S // U_STR_S = "abc"
)
```

### iota
iota 可以认为是一个可以被编译器修改的常量, 可以被用作枚举值,在每一个const关键字出现时，被重置为0，然后再下一个const出现之前，每出现一次iota，其所代表的数字会自动增加1。

````go
const (
    COLOR_RED = iota // 0
    COLOR_GREEN // 1
    COLOR_BLUE // 2
)
````

### 常见的iota示例

- 使用`_`跳过某些值

```go
const (
    COLOR_RED = iota // 0
    _
    COLOR_GREEN // 2
    COLOR_BLUE // 3
)
```

- iota声明中间插队
```go
const (
    COLOR_RED = iota // 0
    Test_num1 = 100 // 100
    COLOR_GREEN = iota // 2
    COLOR_BLUE // 3
)
```

- 定义数量级
```go
const (
    _  = iota
    KB = 1 << (10 * iota)
    MB = 1 << (10 * iota)
    GB = 1 << (10 * iota)
    TB = 1 << (10 * iota)
    PB = 1 << (10 * iota)
)
```