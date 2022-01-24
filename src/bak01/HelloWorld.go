package main

import "fmt"

func main() {
	fmt.Println("hello "+"golang")
	var bv bool =true
	fmt.Println(bv);
	var astr string="string test"
	fmt.Println(astr)
    var f float64=24.98
    fmt.Println(f)

    const  cstr string="HELLO GOLANG"
    fmt.Println(cstr)

	const (
		Unknown = 0
		Female = 1
		Male = 2
	)
    fmt.Println(Female)



}
