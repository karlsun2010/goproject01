package main

import  (
	"fmt"
	)
func main() {
	a :=[3]int{4,67,90}
	var b [5]int
	var c [3]int
	//a=b  // 数组长度是类型的一部分，因此[5]int 和 [3]int 不是相同类型，不能互相赋值
	c=a
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}
