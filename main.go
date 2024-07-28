package main

import "fmt"

// fmt.Println <= print할 때 사용

func canIDrink(age int) bool {
	switch {
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 50:
		return false
	}
	return false
}

func main() {
	// name := "yohan"과 var name string = "yohan"은 같은 코드이다.
	// name := "yohan" // go가 type을 알아서 결정
	fmt.Println(canIDrink(18))
	a := 10
	fmt.Println(a)
	dks := "안요한"
	fmt.Println(dks)
}
