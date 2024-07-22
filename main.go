package main

import (
	"fmt"
	"strings"
)

//fmt.Printl <= print할 때 사용

func superAdd(number ...int) int {
	defer fmt.Println("I'm done") //특정 함수를 현재 함수가 종료되기 직전에 실행
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	// name := "yohan"과 var name String = "yohan"은 같은 코드이다.
	//name := "yohan" //go가 type을 알아서 결정
	totalLength, up := superAdd("yohan")
	fmt.Println(totalLength, up)
}
