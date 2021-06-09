package main

import (
	"fmt"
)

func main() {
	s := "abcd"
	// s[0] = 'c' // 报错：cannot assign to s[0]

	bs := []byte(s)

	bs[1] = 'B'
	s1 :=  string(bs)
	fmt.Println(string(bs), &s, &bs, &s1)

	u := "电脑"
	us := []rune(u)

	us[1] = '话'
	fmt.Println(string(us))
}