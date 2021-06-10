## Channel(管道)
`channel`用来在`goroutine`之间来传递消息。

> Go的设计思路：不要通过共享内存来通信，而要通过通信来实现共享内存。

声明的格式是：
```go
var 变量名 chan 类型
```

```go

var ch0 chan int
var ch1 chan string
var ch2 chan map[string]string

type stu struct{}

var ch3 chan stu
var ch4 chan *stu

```

说明：
> a. 类似unix中管道（pipe）
> 
> b. 先进先出
> 
> c. 线程安全，多个goroutine同时访问，不需要加锁
> 
> d. channel是有类型的，一个整数的channel只能存放整数

### 初始化

通道是引用类型，通道类型的空值是nil。

```go
var ch chan int // <nil>
var ch0 chan int = make(chan int) // 无缓冲
var ch1 chan int = make(chan int, 10) // 有缓冲
```

### channel操作
通道有发送（send）、接收(receive）和关闭（close）三种操作。

* 将一个值发送到通道中

```go
ch <- 10 // 把10发送到ch中
```
* 从一个通道中接收值。

```go
x := <- ch // 从ch中接收值并赋值给变量x
<-ch       // 从ch中接收值，忽略结果
```

*调用内置的close函数来关闭通道

```go
close(ch)
```
只有在通知接收方goroutine所有的数据都发送完毕的时候才需要关闭通道。通道是可以被垃圾回收机制回收的，它和关闭文件是不一样的，在结束操作之后关闭文件是必须要做的，但关闭通道不是必须的。

需要注意的是：
```go
1.对一个关闭的通道再发送值就会导致panic。
2.对一个关闭的通道进行接收会一直获取值直到通道为空。
3.对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
4.关闭一个已经关闭的通道会导致panic。
```


### 深入理解
无缓冲的与有缓冲channel有着重大差别，一个是同步的 一个是非同步的。

#### 无缓冲管道
无缓冲的通道又称为同步通道：通信是同步，在有接受者接收数据之前，发送不会结束。
无缓冲的通道没有空间来保存数据，要求发送/接收操作在对方准备好之前是阻塞的。

```go
func main() {
    ch := make(chan int)
    ch <- 10
    fmt.Println("发送成功")
}
```
输出：
```go
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
```

> 无缓冲的通道只有在有人接收值的时候才能发送值

启动一个协程来接收：
```go
func recv(c chan int) {
    ret := <-c
    fmt.Println("接收成功", ret)
}
func main() {
    ch := make(chan int)
    go recv(ch) // 启用goroutine从通道接收值
    ch <- 10
    fmt.Println("发送成功")
}
```

#### 有缓冲通道
有缓冲通道又称为异步通道，通过判断缓冲区来决定是否阻塞。如果缓冲区已满，发送被阻塞;缓冲区为空，接收被阻塞。

```go
func main() {
    ch := make(chan int, 5) // 创建一个容量为5的有缓冲区通道
    ch <- 10
    fmt.Println("发送成功")
}
```
异步 `channel` 可减少排队阻塞，具备更高的效率。通常使用指针来管道通讯，规避大对象拷贝，将多个元素打包，减小缓冲区大小等

```go
package main

import (
	"fmt"
)

func main() {
	data := make(chan int, 3) // 缓冲区可以存储 3 个元素
	exit := make(chan bool)

	data <- 1 // 在缓冲区未满前，不会阻塞。
	data <- 2
	data <- 3

	go func() {
		for d := range data { // 在缓冲区未空前，不会阻塞。
			fmt.Println(d)
		}

		exit <- true
	}()

	data <- 4 // 如果缓冲区已满，阻塞。
	data <- 5
	close(data)

	<-exit
}
```

输出：
```go
1
2
3
4
5
```

#### select
如果需要同时处理多个 channel，可使用 select 语句。它随机选择一个可用 channel 做收发操作，或执行 default case。

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	a, b := make(chan int, 3), make(chan int)

	go func() {
		v, ok, s := 0, false, ""

		for {
			select {
			case v, ok = <-a:
				s = "a"
			case v, ok = <-b:
				s = "b"
			}

			if ok {
				fmt.Println(s, v)
			} else {
				os.Exit(0)
			}
		}
	}()

	for i := 0; i < 5; i++ {
		select { // 随机选择可  channel，接收数据。
		case a <- i:
		case b <- i:
		}
	}

	close(a)
	select {} // 没有可用 channel，阻塞 main goroutine。
}
```

输出：
```go
// 每次运行输出结果都不同
b 3
a 0
a 1
a 2
b 4
```

实现超时 (timeout)：

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	w := make(chan bool)
	c := make(chan int, 2)

	go func() {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-time.After(time.Second * 3):
			fmt.Println("timeout.")
		}

		w <- true
	}()

	// c <- 1 // 注释掉，引发 timeout。
	<-w
}
```