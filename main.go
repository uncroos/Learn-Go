package main

import "fmt"

//fmt.Printl <= print할 때 사용

func multiply(a int, b int) int {
	return a * b
}

func main() {
	// name := "yohan"과 var name String = "yohan"은 같은 코드이다.
	//name := "yohan" //go가 type을 알아서 결정
	fmt.Println(multiply(3, 4))
}
