package main

import "fmt"

//fmt.Printl <= print할 때 사용

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	// name := "yohan"과 var name String = "yohan"은 같은 코드이다.
	//name := "yohan" //go가 type을 알아서 결정
	repeatMe("yohan", "nico", "mike", "pussto")
}
