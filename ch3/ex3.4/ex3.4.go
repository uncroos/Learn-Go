package main

import "fmt"

func main(){
	srt := "Hello\tGo\t\tWorld\n\"GO\"is Awesome!\n" // 문자열

	fmt.Print(srt)
	fmt.Printf("%s", srt)
	fmt.Printf("%q", srt)
}