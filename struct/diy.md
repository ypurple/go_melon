## 自定义类型
Go语言支持我们自定义类型，比如刚刚上一节的结构体类型，就是我们自定义的类型，这也是比较常用的自定义类型的方法。

另外一个自定义类型的方法是基于一个已有的类型，就是基于一个现有的类型创造新的类型，这种也是使用type关键字。

```go
type Duration int64
```
在使用time这个包的时候，类型time.Duration是基于int64这个基本类型创建的新类型，来表示时间的间隔。

```
type Duration int64
var i Duration = 100
var j int64 = 100
```
i 和 j 都是基于int64创建的，但是本质上，他们并不是同一种类型，所以对于Go这种强类型语言，他们是不能相互赋值的。

```go
type Duration int64
var dur Duration
dur=int64(100) // panic: cannot use int64(100) (type int64) as type Duration in assignment
fmt.Println(dur)
```

自定义类型是Go灵活的地方，可以使用自定义的类型做很多事情，比如添加方法，比如可以更明确的表示业务的含义等等

```go
package main

import (
	"fmt"
)
type Callback func(in string)

func cb1(in string) {
	fmt.Println("cb1 :" + in)
}
func cb2(in string) {
	fmt.Println("cb2 :" + in)
}

func Uprint(in string, callback Callback) {
	callback(in)
}

func main() {
	var in = "test callback func"
	Uprint(in, cb1)
	Uprint(in, cb2)
}
```