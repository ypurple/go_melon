## 基本类型介绍
Golang 更明确的数字类型命名，支持 Unicode，支持常用数据结构。

- 基本类型：

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
|string                |       |   ""  | UTF-8 字符串

- 高级类型：

|类型	|长度(字节) |	默认值	|说明
|-------|--------- |------------|-----
|struct                |       |   | 值类型
|slice         |       |   nil | 引用类型
|map           |       |   nil | 引用类型
|channel       |       |       nil | 引用类型
|interface     |       |       nil | 接口
|function      |       |       nil | 函数
