package main

import "fmt"

//go语言中 数组是值类型，不是引用类型
func main() {
	a:= [...]string{"zs","ls","ww","zl"}
	b := a
	b[2] = "ww222"
	c :=a
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
