package main

import "fmt"

func main() {
	var rows = 5
	for i := 1; i <= rows; i++ {
		diff:= rows-1
		k:=i
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", k)
			k = k + diff
			diff = diff-1
		}
		fmt.Println()
	}
}
