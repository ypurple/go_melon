## Go的包管理
从`golang 1.11`开始，`golang`提供`go mod`命令来管理包，使得使用外部包变得特别轻松简单


## go mod 简介
```go
$ go help mod

Go mod provides access to operations on modules.

Note that support for modules is built into all the go commands,
not just 'go mod'. For example, day-to-day adding, removing, upgrading,
and downgrading of dependencies should be done using 'go get'.
See 'go help modules' for an overview of module functionality.

Usage:

    go mod <command> [arguments]

The commands are:

    download    download modules to local cache
    edit        edit go.mod from tools or scripts
    graph       print module requirement graph
    init        initialize new module in current directory
    tidy        add missing and remove unused modules
    vendor      make vendored copy of dependencies
    verify      verify dependencies have expected content
    why         explain why packages or modules are needed

Use "go help mod <command>" for more information about a command.
```

> go mod 有以下命令:

|命令 |	说明
|---- |-----
|download	| 下载依赖包
|edit	    | 编辑go.mod
|graph	    | 打印模块依赖图
|init	    | 在当前目录初始化mod
|tidy	    | 拉取缺少的模块，移除不用的模块
|vendor	    | 将依赖复制到vendor下
|verify	    | 验证依赖是否正确
|why	    | 解释为什么需要依赖

比较常用的是 `init`, `tidy`, `edit`

## 初始化项目

```go
$ mkdir domo
$ cd domo
$ go mod init domo
```

查看一下 go.mod文件：
```go
$ cat go.mod
module demo

go 1.14
```

`go.mod`文件一旦创建后，它的内容将会被`go toolchain`全面掌控。
`go toolchain`会在各类命令执行时，比如`go get`、`go build`、`go mod`等修改和维护`go.mod`文件

## 添加依赖

```go
package main

import "fmt"
import "rsc.io/quote"

func main() {
    fmt.Println(quote.Go())
}
```

执行 `go run main.go` 运行代码会发现 `go mod` 会自动查找依赖自动下载:
```go
module demo

go 1.14

require (
    rsc.io/quote v1.5.2 // indirect
)
```

* `go module` 安装 `package`的原則是先拉最新的`release tag`，若无tag则拉最新的`commit`;
* go 会自动生成一个 go.sum 文件来记录 `dependency tree`

> go.mod说明
 
```go
- module 语句指定包的名字（路径）
- require 语句指定的依赖项模块
- replace 语句可以替换依赖项模块
- exclude 语句可以忽略依赖项模块
```

> 1、当依赖包有更新时，删除`go.mod`中对应包的行执行 `go build`或者直接执行`go get -u need-upgrade-package`，即可加载最新包

> 2、如果替换远程包为本地包使用`replace rsc.io/quote ../quote` 即可


