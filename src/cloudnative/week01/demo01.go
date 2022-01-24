package main
/*
1 字符串遍历
2 数组定义  遍历，并修改数组的值
*/
import (
	"fmt"
)

func main() {
	var str string = "hello world"
	fmt.Println(str)
	for index, char := range str {
		fmt.Println(index)
		fmt.Println(string(char))
		//fmt.Println(char)
		//fmt.Println(string(char))


	}
		var myArray =[5]string{"I","am","stupid","and","weak"}
		fmt.Println(myArray)
		for _,value := range myArray{
			fmt.Println(value)
		}
	for index, _ := range myArray{
		if(index==2){
			myArray[index]="smart"
		}else if(index ==4){
			myArray[index]="strong"
		}
	}
	fmt.Println(myArray)
}
