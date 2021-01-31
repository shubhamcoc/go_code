package main

import (
	"fmt"
)

func main() {
	arr := []int{6, 10, 14, 5, 8, 3, 12, 1}
	fmt.Println(arr)
	result := selectionSort(arr)
	fmt.Println(result)
}

func selectionSort(A []int) []int {
	size := len(A)
	for i := 0; i < size-1; i++ {
		min := i
		for j := i+1; j < size; j++ {
			if A[j] < A[min] {
				min = j
			}
		}
		
		if min != i {
			A[i], A[min] = A[min], A[i]
		}
	}
	
	return A
}
