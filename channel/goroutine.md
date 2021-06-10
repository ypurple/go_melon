## goroutine（协程）和channel（管道）
`goroutine` 和 `channel`往往在一起，比较好理解，`channel`属于引用类型，所以一起放在本节

Go中，`goroutine`和`channel`是并发编程的两大基石，`goroutine`用来执行并发任务，`channel`用来在`goroutine`之间来传递消息。

### goroutine(协程)
`goroutine`协程是一种用户态的轻量级线程，协程的调度完全由用户控制，协程间切换只需要保存任务的上下文，没有内核的开销。

> Goroutine 非常轻量
> 
> 上下文切换代价小
> 
> 内存占用少：线程栈空间通常是 2M，`Goroutine`栈空间最小 2K；轻松支持10w 级别的 `Goroutine` 运行
> 
> Go内部实现了G-P-M模型 调度器

### 协程使用
Go语言中使用`goroutine`非常简单，只需要在调用函数的时候在前面加上go关键字，就可以为一个函数创建一个goroutine。

```go
func hello() {
    fmt.Println("Hello Goroutine!")
}
func main() {
    go hello()
    fmt.Println("main goroutine done!")
}
```
执行结果只打印了`main goroutine done!`，并没有打印`Hello Goroutine!`。为什么呢？

原来当main()函数返回的时候进程就结束了，`goroutine`还没来得及执行，就一同结束；我们可以加一个`sleep` 看效果
```go
func main() {
    go hello() // 启动另外一个goroutine去执行hello函数
    fmt.Println("main goroutine done!")
    time.Sleep(2 * time.Second)
}
```

输出：
```go
main goroutine done!
Hello Goroutine!
```

为什么会先打印`main goroutine done!`?

是因为我们在创建新的`goroutine`的时候需要花费一些时间，而此时`main`函数所在的`goroutine`是继续执行的。

### 协程同步
`WaitGroup`用途：一直等到所有的`goroutine`执行完成，并且阻塞主线程的执行，直到所有的`goroutine`执行完成。
三个方法的作用。
```
Add: 添加或者减少等待goroutine的数量

Done: 相当于Add(-1)

Wait: 执行阻塞，直到所有的WaitGroup数量变成0
```

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    wg := sync.WaitGroup{}

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go calc(&wg, i)
    }

    wg.Wait()
    fmt.Println("all goroutine finish")
}
func calc(w *sync.WaitGroup, i int) {

    fmt.Println("calc:", i)
    time.Sleep(time.Second)
    w.Done()
}
```
输出：
```
calc: 1
calc: 5
calc: 7
calc: 8
calc: 2
calc: 4
calc: 6
calc: 0
calc: 3
calc: 9
all goroutine finish

```