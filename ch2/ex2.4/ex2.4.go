package main

import "fmt"

func main(){
	a := 3	// int 형
	var b float64 = 3.5 // float64형

	var c int = int(b)	//float64형에서 int형으로 변환
	d := float64(a*c)	//int형 에서 float64형으로 변환

	var e int64 = 7		
	f := int64(d)*e		// float64형에서 int64형으로 변환

	var g int = int(b * 3)	//float64형에서 int형으로 변환
	var h int = int(b)*3	//float64형에서 int형으로 변환. g와 값이 다름
	fmt.Println(g,h,f)

}