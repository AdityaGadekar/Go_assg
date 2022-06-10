package main

import "fmt"

func bubble_sort(arr []int, size int) []int {
	for i:=0;i<size ;i++{
		for j:=0; j<size-1;j++{
			if arr[j]> arr[j+1]{
				arr[j], arr[j+1]= arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	fmt.Printf("Enter size of array\n")
	var size int
	fmt.Scanln(&size)
	// var arr = []int{1,2,3,4,5}
	arr := make([]int, size)
	fmt.Println("Enter Elements ")
	for i:=0; i<size; i++ {
		fmt.Scanln(&arr[i])
	}
	fmt.Println("Array : ", arr)
	fmt.Println(bubble_sort(arr,size))
}
