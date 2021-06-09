### 基本类型介绍
Golang的引用类型包括 `slice`、`map` 和 `channel`。它们有复杂的内部结构，除了申请内存外，还需要初始化相关属性。
> 变量存储的是一个地址，这个地址存储最终的值。
>
> 内存通常在堆上分配。通过GC回收。
>
>获取指针类型所指向的值，使用：" * " 取值符号 。比如：var *p int, 使用*p获取p指向的值指针


- 引用类型：

|类型	|默认值	|说明
|-------|------------|-----
|slice		| nil | 引用类型
|map		| nil | 引用类型
|channel	| nil | 引用类型

#### new 和 make

- 内置函数 new 计算类型大小，为其分配零值内存，返回指针。 
- make 会被编译器翻译 成具体的创建函数，由其分配内存和初始化成员结构，返回对象而非指针。

```go
package main

func main() {
	a := []int{0, 0, 0} // 提供初始化表达式。
	a[1] = 10

	b := make([]int, 3) // make slice
	b[1] = 10

	c := new([]int)
	c[1] = 10 // ./main.go:11:3: invalid operation: c[1] (type *[]int does not support indexing)

}
```

new和make的区别：
```go
make 用来创建`map`、`slice`、`channel`等引用类型 

new 用来创建值类型(如`array`, `struct`) 和 用户定义的类型

new 和 make 均是用于分配内存

```


