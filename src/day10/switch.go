package main

import "fmt"

func main() {
	var flag int = 4
	fmt.Println(flag)
	switch flag:=9;flag {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4,5,6,7,8:
		fmt.Println("大于 等于 four ")
	default:
		fmt.Println("error number")
	}
}
