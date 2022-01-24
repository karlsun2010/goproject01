package main

import "fmt"
func main() {
	var personMaps map[string]int
	fmt.Println(len(personMaps))
	//
	if personMaps== nil {
		fmt.Println("map personMaps is nil")
		personMaps = make(map[string]int)
	}
	personMaps["zs"]=1500
	personMaps["ls"]=1700
	personMaps["ww"]=2300
	fmt.Println(len(personMaps))
	fmt.Println("personMaps is:",personMaps)

}
