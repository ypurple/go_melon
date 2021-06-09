package main

import (
	"fmt"
	"strings"
)

func prefix(){
	str := "hello world"
	res0 := strings.HasPrefix(str, "http://")
	res1 := strings.HasPrefix(str, "hello")
	fmt.Printf("res0 is %v\n", res0)
	fmt.Printf("res1 is %v\n", res1)
}

func suffix() {
	str := "hello world"
	res0 := strings.HasSuffix(str, "http://")
	res1 := strings.HasSuffix(str, "world")
	fmt.Printf("res0 is %v\n", res0)
	fmt.Printf("res1 is %v\n", res1)
}

func index() {
	str := "hello world"
	res0 := strings.Index(str, "o")
	res1 := strings.Index(str, "i")
	fmt.Printf("res0 is %v\n", res0)
	fmt.Printf("res1 is %v\n", res1)
}

func LastIndex() {
	str := "hello world"
	res0 := strings.LastIndex(str, "o")
	res1 := strings.LastIndex(str, "i")
	fmt.Printf("res0 is %v\n", res0)
	fmt.Printf("res1 is %v\n", res1)
}

func Replace() {
	str := "hello world world"
	res0 := strings.Replace(str, "world", "golang", 2)
	res1 := strings.Replace(str, "world", "golang", 1)
	//trings.Replace("原字符串", "被替换的内容", "替换的内容", 替换次数)
	fmt.Printf("res0 is %v\n", res0)
	fmt.Printf("res1 is %v\n", res1)
}

func Count() {
	str := "hello world world"
	countTime0 := strings.Count(str, "o")
	countTime1 := strings.Count(str, "i")
	fmt.Printf("countTime0 is %v\n", countTime0)
	fmt.Printf("countTime1 is %v\n", countTime1)
}

func Repeat() {
	str := "hello world "
	res0 := strings.Repeat(str, 0)
	res1 := strings.Repeat(str, 1)
	res2 := strings.Repeat(str, 2)
	// strings.Repeat("原字符串", 重复次数)
	fmt.Printf("res0 is %v\n", res0)
	fmt.Printf("res1 is %v\n", res1)
	fmt.Printf("res2 is %v\n", res2)
}

func ToUpper() {
	str := "hello world "
	res := strings.ToUpper(str)

	fmt.Printf("res is %v\n", res)
}


func ToLower() {
	str := "HELLO WORLD "
	res := strings.ToLower(str)

	fmt.Printf("res is %v\n", res)
}

func TrimSpace() {
	str := "     hello world     "
	res := strings.TrimSpace(str)

	fmt.Printf("res is %v\n", res)
}


func Trim() {
	str := "hi , hello world , hi"
	res := strings.Trim(str, "hi")

	fmt.Printf("res is %v\n", res)
}


func TrimLeft() {
	str := "hi , hello world , hi"
	res := strings.TrimLeft(str, "hi")

	fmt.Printf("res is %v\n", res)
}


func TrimRight() {
	str := "hi , hello world , hi"
	res := strings.TrimRight(str, "hi")

	fmt.Printf("res is %v\n", res)
}


func Fields() {
	str := "hello world ，hello golang"
	res := strings.Fields(str)

	fmt.Printf("res is %v\n", res)
}

func Join() {
	str := []string{"hello", "world", "hello", "golang"}
	res := strings.Join(str, "++")

	fmt.Printf("res is %v\n", res)
	/*
		num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
		res1 := strings.Join(num, "++")
		//  cannot use num (type []int) as type []string in argument to strings.Join
		fmt.Println(res1)
	*/
}


func Split() {
	str := "hello world ，hello golang"
	res := strings.Split(str, "o")

	fmt.Printf("res is %v\n", res)
}

func main() {

	prefix()

	suffix()

	index()

	LastIndex()

	Replace()

	Count()

	Repeat()

	ToUpper()

	ToLower()

	TrimSpace()

	Trim()

	TrimLeft()

	TrimRight()

	Fields()

	Join()

	Split()
}