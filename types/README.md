# go 基础
这小节我们将要介绍如何定义变量、常量、Go内置类型以及Go程序设计中的一些技巧

### Golang文件名
```go
所有的go源码都是以 ".go" 结尾
```

### Go语言命名
> Go的函数、变量、常量、自定义类型、包(package)的命名方式遵循以下规则：
 
```go
 1）首字符可以是任意的Unicode字符或者下划线
 2）剩余字符可以是Unicode字符、下划线、数字
 3）字符长度不限
```

> Go有25个关键字:

```go
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

> Go还有37个保留字:

```go
Constants:    true  false  iota  nil

Types:    int  int8  int16  int32  int64  
          uint  uint8  uint16  uint32  uint64  uintptr
          float32  float64  complex128  complex64
          bool  byte  rune  string  error

Functions:   
    append  		-- 用来追加元素到数组、slice中,返回修改后的数组、slice
    close   		-- 主要用来关闭channel
    delete    		-- 从map中删除key对应的value
    panic    		-- 停止常规的goroutine  （panic和recover：用来做错误处理）
    recover 		-- 允许程序定义goroutine的panic动作
    make    		-- 用来分配内存，返回Type本身(只能应用于slice, map, channel)
    new        		-- 用来分配内存，主要用来分配值类型，比如int、struct。返回指向Type的指针
    cap        		-- capacity是容量的意思，用于返回某个类型的最大容量（只能用于切片和 map）
    copy    		-- 用于复制和连接slice，返回复制的数目
    len        		-- 来求长度，比如string、array、slice、map、channel ，返回长度
    complex             -- 创建复数
    imag    		-- 返回complex的实部   （complex、real imag：用于创建和操作复数）
    real    		-- 返回complex的虚部
```