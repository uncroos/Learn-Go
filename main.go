package main

import "fmt"

//fmt.Printl <= print할 때 사용

func superAdd(numbers ...int) int {
	total := 0
	for index, number := range numbers {
		fmt.Println(index, number)
	}
	return 1
}

func main() {
	// name := "yohan"과 var name String = "yohan"은 같은 코드이다.
	//name := "yohan" //go가 type을 알아서 결정
	superAdd(1, 2, 3, 4, 5)
}
