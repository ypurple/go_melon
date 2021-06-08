# hello word
我们的第一个程序将打印传说中的 "hello world"消息，下面是完整的程序代码。

```go
package main

import "fmt"

func main() {
    fmt.Println("hello world")
}
```

要运行这个程序，将这些代码放到 `hello-world.go` 中并且使用`go run` 命令
```go
$ go run hello-world.go
hello world
```

如果我们想将我们的程序编译成二进制文件。我们可以通过`go build`命来达到目的，然后我们可以直接运行这个二进制文件
```go
$ go build hello-world.go
$ ls
hello-world	hello-world.go

$ ./hello-world
hello world

```

## 详解
* Go程序是通过package来组织的，`package main`是当前文件属于哪个包。
* 每一个可独立运行的Go程序，必定包含一个`package main`，`main`包中必定包含一个入口函数`main`，这个函数既没有参数，也没有返回值。
* 为了打印 `Hello world`，我们调用了一个函数`Printf`，这个函数来自于fmt包，所以需要导入fmt包：`import "fmt"`
* 包的概念和Python中的package类似，好处：模块化（能够把你的程序分成多个模块)和可重用性（每个模块都能被其它应用程序反复使用）

