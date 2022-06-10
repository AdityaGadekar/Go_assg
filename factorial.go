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

	fmt.Println("\nFactorial by recursion =" ,(factorial(num)))

}

func factorial(n int)int{
	if n == 1{
		return 1
	} 
	return n*factorial(n-1)

}
