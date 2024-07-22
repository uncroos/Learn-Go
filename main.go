package main

import "strings"

//fmt.Printl <= print할 때 사용

func lebAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
	// name := "yohan"과 var name String = "yohan"은 같은 코드이다.
	//name := "yohan" //go가 type을 알아서 결정
	totalLenght, upperName := lebAndUpper("yohan")
}
