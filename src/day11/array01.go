package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a [3]int
	a[0]=2
	a[1]=5
	fmt.Println(a)
    b:=[5]int{3,56,78,343,23}
    fmt.Println(b)
    c :=[...]int{45,67,23}
    fmt.Println(c)
    fmt.Println("array c length is "+strconv.Itoa(len(c))) // string 转 int

}
