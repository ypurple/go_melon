## 基本类型介绍
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
