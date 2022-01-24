package main


import (
	"fmt"
	"strings"
)


func main() {
	fmt.Println("hello go")

	var bv  bool=true
fmt.Println(&bv)
	fmt.Println(bv)

     str :="这里是 www\n.runoob\n.com"
     fmt.Println(str)

     str=strings.Replace(str," ","",-1)
     str=strings.Replace(str,"\n","",-1)
     fmt.Println(str)
     fmt.Println(&str)

	const   LENGTH  = 25
	fmt.Println(LENGTH)

	var a int =10
	if a < 20 {
		fmt.Println("a =",a)
	}

}

