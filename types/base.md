### 基本类型介绍
Golang 更明确的数字类型命名，支持 Unicode，支持常用数据结构。

|类型	|长度(字节) |	默认值	|说明
|-------|--------- |------------|-----
|bool	| 1         |false      |
|byte	| 1	        |0	        |   uint8
|rune	| 4	        |0	        |   Unicode Code Point, int32
|int, uint    |	4或8 |	0	    |   32 或 64 位
|int8, uint8  |	1	|0	        |-128 ~ 127, 0 ~ 255，byte是uint8 的别名
|int16, uint16|	2	|0	|-32768 ~ 32767, 0 ~ 65535
|int32, uint32|	4	|0	|-21亿~ 21亿, 0 ~ 42亿，rune是int32 的别名
|int64, uint64|	8	|0  | |
|float32    | 4	    |0.0| |
|float64	| 8	    |0.0| |
|complex64	| 8		|   | |
|complex128	| 16	|	|   |
|uintptr	| 4或8	|	| 以存储指针的 uint32 或 uint64 整数
|array		|       |	| 值类型
|struct		|       |   | 值类型
|string		|       |   ""  | UTF-8 字符串
|slice		|       |   nil | 引用类型
|map		|       |   nil | 引用类型
|channel	|       |	nil | 引用类型
|interface	|       |	nil | 接口
|function	|       |	nil | 函数

#### 整型
整型分为以下两个大类： 
- 按长度分为：int8、int16、int32、int64
- 对应的无符号整型：uint8、uint16、uint32、uint64
- 其中，uint8就是我们熟知的byte型，int16对应C语言中的short型，int64对应C语言中的long型。

#### 浮点型
Go语言支持两种浮点型数：`float32` 和 `float64`
- float32 的浮点数的最大范围约为3.4e38，可以使用常量定义：math.MaxFloat32。 
- float64 的浮点数的最大范围约为 1.8e308，可以使用一个常量定义：math.MaxFloat64。

#### 复数

`complex64`和`complex128`
复数有实部和虚部，complex64的实部和虚部为32位，complex128的实部和虚部为64位。

复数使用 re+imI 来表示，其中 re 代表实数部分，im 代表虚数部分，I 代表根号负 1。
```go
var c1 complex64 = 5 + 10i
fmt.Printf("The value is: %v", c1)
// 输出： 5 + 10i

```
函数 real(c) 和 imag(c) 可以分别获得相应的实数和虚数部分。

#### 布尔值
Go语言中以bool类型进行声明布尔型数据，布尔型数据只有true（真）和false（假）两个值。

注意：
- 布尔类型变量的默认值为false。
- Go 语言中不允许将整型强制转换为布尔型.
- 布尔型无法参与数值运算，也无法与其他类型进行转换。

#### 字符串

Go 语言里的字符串的内部实现使用UTF-8编码。 字符串的值为双引号(")中的内容，可以在Go语言的源码中直接添加非ASCII码字符
```go
s1 := "hello"
s2 := "你好"
```

> • 默认值是空字符串 ""。

> • 用索引号访问某字节，如 s[i]

> • 不能用序号获取字节元素指针，&s[i] 非法

> • 不可变类型，无法修改字节数组

> • 字节数组尾部不包含 NULL

```go
package main

func main() {
	s := "abc"
	println(s[0] == '\x61', s[1] == 'b', s[2] == 0x63)
}

```
输出结果:
```go
true true true
```

##### 字符串转义符
Go 语言的字符串常见转义符包含回车、换行、单双引号、制表符等

|转义	| 含义
|------ |------
|\r	    | 回车符（返回行首）
|\n	    | 换行符（直接跳到下一行的同列位置）
|\t	    | 制表符
|\'	    | 单引号
|\"	    | 双引号
|\	    | 反斜杠


##### 多行字符串
多行字符串时，使用反引号字符：
```go
package main

func main() {
	s1 := `轻轻的我走了，
        正如我轻轻的来；
        我轻轻的招手，
        作别西天的云彩。
    `
    fmt.Println(s1)
}

```
反引号间换行将被作为字符串中的换行，但是所有的转义字符均无效，文本将会原样输出

> 连接跨行字符串时，"+" 必须在上一行末尾，否则导致编译错误:

```go
package main

import (
	"fmt"
)

func main() {
	s := "Hello, " +
		"World!"
	// s2 := "Hello, "
	// +"World!" 
    //./main.go:11:2: invalid operation: + untyped string

	fmt.Println(s)
}
```

##### 字符串的常用操作

|方法	 |   介绍
|------- |---------
|len(str)	    |求长度
|+或fmt.Sprintf	|拼接字符串
|strings.Split	|分割
|strings.Contains	|   判断是否包含
|strings.HasPrefix, strings.HasSuffix	|前缀/后缀判断
|strings.Index(), strings.LastIndex()	|子串出现的位置
|strings.Join(a[]string, sep string)	|join操作
|strings.Replace	|替换
|strings.Count      |子串出现的次数
|strings.Repeat     |字符串重复次数
|strings.ToUpper， strings.ToLower| 转成大/小写
|strings.TrimSpace，strings.Trim  |去掉 str 首尾的空格/某个字符

如：
```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello world"
	res0 := strings.HasPrefix(str, "http://")
	res1 := strings.HasPrefix(str, "hello")
	fmt.Printf("res0 is %v\n", res0)
	fmt.Printf("res1 is %v\n", res1)
}
```
输出：
```go
res0 is false
res1 is true
```

##### byte和rune类型

![img.png](img.png)

组成每个字符串的元素叫做“字符”，可以通过遍历或者单个获取字符串元素获得字符。 字符用单引号（’）包裹起来
```go
var a := '中'
var b := 'x'
```

Go 语言的字符有以下两种：
- uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
- rune类型，代表一个 UTF-8字符，实际是一个int32

当需要处理中文、日文或者其他复合字符时，则需要用到rune类型，因为UTF8编码下一个中文汉字由3~4个字节组成
```go
// 遍历字符串
func scanString() {
    s := "hello.cn你好"
    for i := 0; i < len(s); i++ { //byte
        fmt.Printf("%v(%c) ", s[i], s[i])
    }
    fmt.Println()
    for _, r := range s { //rune
        fmt.Printf("%v(%c) ", r, r)
    }
    fmt.Println()
}
```

输出：
```go
104(h) 101(e) 108(l) 108(l) 111(o) 46(.) 99(c) 110(n) 228(ä) 189(½) 160( ) 229(å) 165(¥) 189(½) 
104(h) 101(e) 108(l) 108(l) 111(o) 46(.) 99(c) 110(n) 20320(你) 22909(好) 
```

##### 修改字符串
修改字符串，先将其转换成[]rune或[]byte，完成后再转换为string。无论哪种转换，都会重新分配内存，并复制字节数组。
```go
package main

func main() {
	s := "abcd"
	bs := []byte(s)

	bs[1] = 'B'
	println(string(bs))

	u := "电脑"
	us := []rune(u)

	us[1] = '话'
	println(string(us))
}
```

输出：
```go
aBcd
电话
```