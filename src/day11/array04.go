package main

import "fmt"

func main() {
	a:= [3][2]string{{"01","zs"},{"02","ls"},{"03","ww"},
	}
	fmt.Println(a)
	var b [3][2]string
	b[0][0]="01"
	b[0][1]="zs"
	fmt.Println(b)
}
