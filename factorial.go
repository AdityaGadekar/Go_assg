package main

import "fmt"

func main(){
	var i, num int
	fact:=1
	fmt.Println("Enter num")
	fmt.Scanf("%d", &num)
	for i=1; i<=num;i++{
		fact =fact*i
	}
	fmt.Printf("Factorial = %d", fact)

}