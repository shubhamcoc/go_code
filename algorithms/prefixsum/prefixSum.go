package main

import (
	"fmt"
)

func main() {
	test := []int{1, 2, 3, 4, 5}
	fmt.Println(test)
	prefixSumArray(test)  // pass by reference in Golang
	fmt.Println(test)
	result := prefixSum(test, 2, 4)
	result1 := prefixSum(test, 1, 3) // reduce the time complexity to O(1)
	fmt.Println(result)
	fmt.Println(result1)
}

func prefixSumArray(arr []int) {
	sum := 0
	for i := 0; i<=len(arr)-1; i++ {
		sum+=arr[i]
		arr[i] = sum
	}
}

func prefixSum(arr []int, l, h int) int {
	
	res := arr[h] - arr[l-1]
	return res
}
