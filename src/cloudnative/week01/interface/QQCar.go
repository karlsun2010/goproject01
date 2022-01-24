package main

import "fmt"

type Car interface {
	run()
	stop()
}

type QQ interface {
	login()
}

type QQCar struct {


}





func (qqCar QQCar)run(){
fmt.Println("QQCare run")

}
func (qqCar QQCar)stop(){

fmt.Println("QQCare stop")
}

func (qqCar QQCar)login(){
	fmt.Println("QQCare login")

}
func main() {
	var car  Car
	car = new(QQCar)
	car.run()
	var qq QQ
	qq=new(QQCar)
	qq.login()

}