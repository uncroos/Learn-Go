package main

import "fmt"

func main(){
	var a int = 10							// a 변수 선언
	var msg string = "Hello Variable"		// msg 변수 선언

	a = 20									// a 값 변경
	msg = "Good Morning"					// msg값 변경
	fmt.Println(msg, a)						// msg와 a 값 출력
}