package main

import "errors"

func  Sqrt(fin float64)( float64, error){
if fin< 0 {
retrun 0,errors.New("test error")

}

}
func Sqrt2(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	// 实现
	return  f
}


func main() {
	
}
