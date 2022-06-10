package main

import (
	"fmt"
)

func string_reverse(s string) string{
	slice := []rune(s)
	len:= len(slice)
	fmt.Println(slice)
	println(slice[5])
	fmt.Println(len)
	for i:=0; i<len/2; i++ {
		slice[len-i-1] , slice[i] = slice[i] , slice[len-i-1]
	}
	return string(slice)
}

func main(){
	var rev string
	fmt.Print("Enter String\n")
	fmt.Scanln(&rev)
	fmt.Println(string_reverse(rev))
	
}