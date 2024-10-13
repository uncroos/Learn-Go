package main

import "fmt"

var g int = 10 //패키기 전역 변수 선언

func main(){
	var m int = 20 //지역 변수 선언

	{	//지역 변수 선언
		var s int = 50
		fmt.Println(m,s,g)
		//중괄호 닫음 => s 지역 변수 사라짐
	}
	
	// m = s + 20 //Error
}