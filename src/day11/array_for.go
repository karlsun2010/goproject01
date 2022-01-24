package main

import "fmt"

func main() {
	a := [...]string{"dfd","dfd","dfd","dfd","dfd"}
	fmt.Println(a)
	for i:=0;i<len(a);i++ {fmt.Println(a[i])
	}

	for i,v := range a{
		fmt.Printf("%d the element of a is %s\n", i, v)
	}
}
