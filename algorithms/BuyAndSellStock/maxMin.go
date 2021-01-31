package main

import (
	"fmt"
	"math"
)

var cal int=0

func main() {
	min := int(math.Inf(-1))
	max := int(math.Inf(1))
	
	fmt.Println(max)
	fmt.Println(min)
	
	arr := []int{45, 67, 89, -10, -2, 55, 34, -19, 20}
	size := len(arr)
	max, min = findMaxMin(arr, 0, size-1, max, min)
	fmt.Println(max, min)
}

func findMaxMin(arr []int, low, high, max, min int) (int, int) {
	cal = cal + 1
	fmt.Println(cal)
	// If only one element is present then
	// compare the max with it to check if max
	// is more or the element present in the array
	// similarly check if min is lower than the
	// element and return the max and min
	if low == high {
		if max < arr[high] {
			max = arr[high]
		}
		if min > arr[low] {
			min = arr[low]
		}
		return max, min
	}
	
	if high - low == 1 {
		if arr[high] > arr[low] {
			if max < arr[high] {
				max = arr[high]
			}
			
			if min > arr[low] {
				min = arr[low]
			}
		} else {
			if max < arr[low] {
				max = arr[low]
			}
			
			if min > arr[high] {
				min = arr[high]
			}
		}
		
		return max, min
	}
	
	mid := (high + low) / 2
	max, min = findMaxMin(arr, low, mid, max, min)
	max, min = findMaxMin(arr, mid+1, high, max,min)
	
	return max, min
}
