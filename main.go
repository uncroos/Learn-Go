package main

import "fmt"

//fmt.Printl <= print할 때 사용

func canIDrink(age int) bool {
	if age < 18 {
		return false
	} else {
		return true
	}
}

func main() {
	// name := "yohan"과 var name String = "yohan"은 같은 코드이다.
	//name := "yohan" //go가 type을 알아서 결정
	fmt.Println(canIDrink(18))
}
