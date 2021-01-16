package main

import (
	"fmt"
)

func main() {
	arr := [16]int{3,6,8,12,14,17,25,29,31,36,42,47,54,67,80,85}
	fmt.Println(arr)
	res := search(arr, 12)
	fmt.Println(res)
	if res >= 0 { 
		fmt.Println(arr[res])
	} else {
		fmt.Println("Element not found")
	}
}

// Binary Search 
func search(a [16]int, key int) int {
	l := 0
	h := len(a)
	var mid int
	
	// loop should run till low value is lower or equal to high value
	for ; l<=h; {
		// calculating the mid of the array
		mid = (l + h)/2
		fmt.Println(mid, a[mid])
		if key == a[mid] {
			return mid
		}
		// If element is greater than key
		// reduce the high value to (mid - 1)
		// If element is smaller than key
		// increase the low value to (mid + 1)
		if key < a[mid] {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
