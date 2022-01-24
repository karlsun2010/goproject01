package main

import "fmt"

//定义结构体Person
type Person struct {
	fristname ,lastname string
	age int
}

// 匿名结构体
var person struct{
	firstname string
	age  int

}

func main() {
	p1 := Person{
		fristname:"zhang",
		lastname:"san",
		 age:51,
	}
	p2 :=Person{"li","si",23}
	fmt.Println("p1==",p1)
	fmt.Println("p2==",p2)
	var p3 Person;
	fmt.Println("p3==",p3)
	fmt.Println("p1.firstname=",p1.fristname,",p1.lastname="+p1.lastname)

}