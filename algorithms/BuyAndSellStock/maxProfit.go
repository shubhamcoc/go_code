package main

import (
	"fmt"
)

func main() {
	var input []int
	input = append(input, 10)
	input = append(input, 3)
	input = append(input, 5)
	input = append(input, 1)
	input = append(input, 2)
	input = append(input, 8)
	fmt.Println(input)
	profit := maxProfit(input)
	fmt.Println(profit)
}

func maxProfit(arr []int) int {
	min := arr[0]
	index := 0
	
	// Find the min in array and store the index
	for i := 1; i < len(arr); i++ {
		if min > arr[i] {
		    min = arr[i]
		    index = i
		}
	}
	
	// Function to find max in array
	f := func(n []int) int{
		max := n[0]
		
		for i := 1; i < len(n); i++ {
			if max < n[i] {
			    max = n[i]
			}
		}
		return max
	}
	
	if index == len(arr) - 1 {
	    return 0
	}else{
	    // Send the rest of the array from the min index
	    maximum := f(arr[index:])
            fmt.Println("maximum is:", maximum)
            fmt.Println("minimum is:", min)
	    return (maximum - min)
	}
}
