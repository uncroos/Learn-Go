package main

import "fmt"

func main(){
	var a int16 = 3456
	var c int8 = int8(a) //int16형 타입에서 int8형으로 타입 변환'

	fmt.Println(a)
	fmt.Println(c)
}