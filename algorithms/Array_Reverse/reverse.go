package main

import (
	"fmt"
)

func main() {
	arr := [9]int{1,2,3,4,5,6,7,8,9}
	
	fmt.Println(arr)
	res := reverseArray(arr)
	fmt.Println(res)
}

func reverseArray(arr [9]int) [9]int {
	size := len(arr) - 1
        // To calculate no. of executions 
	cal := 0
        // Work Around for running loop with 2 variable
        // as normal syntax is not working.
	for i, j := 0, size; i<j; i,j = i+1, j-1 {
		cal = cal + 1
		arr[i], arr[j] = arr[j], arr[i]
	}
	fmt.Println(cal)
	return arr
}
