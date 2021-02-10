package main

import (
	"fmt"
)

// Returns Sum of array[0...indexed]. This function
// assume that the array is preprocessed and partial
// sum of the array elements stored in the BITtree[]
func getSum(BITtree []int, index int) int {
	sum := 0
	
	// index in the BIT tree is 1 more than the array
	index = index + 1
	
	// Traverse ancestors of BITTree[index]
	for i := index; i>0; {
		// Add the element of the index into the sum
		sum += BITtree[i]
		
		// This is break the index into the 2^ elements
		// i = 10 will be 2^3 + 2^1
		// index will become 10 - 2 = 8
		// i = 12 will be 2^3 + 2^2
		// index will become 12 - 4 = 8
		// i = 6 will be 2^2 + 2^1
		// index will become 6 - 2 = 4
		// i = 13 will be 2^3 + 2^2 + 2^0
		// index will become 13 - 1 = 12
		i -= i & (-i)
	}
	
	return sum
}

// Update Node in BIT at a given index in BIT
// val will be added to BIT[i] and all its ancestors in tree
func updateBIT(BITtree []int, size, index, val int) {
	index = index + 1
	
	for i := index; i <= size; {
		BITtree[i] += val
		
		i += i & (-i)
	}
}

// Construct and return an array of size n
func construct(arr []int, n int) []int {
	BIT := make([]int, n+1)
	
	for i := 0 ; i < n; i++ {
		updateBIT(BIT, n, i, arr[i])
	}
	
	return BIT
}

// query function will return the sum between the 2 indexed
func query(BIT []int, left, right int) int {
	diff := 0
	
	diff = getSum(BIT, right) - getSum(BIT, left-1)
	return diff
}


func main() {
	input := []int{1,2,3,4,5,6,7,8,9,10,11,12}
	
	bit := construct(input, len(input))
	
	fmt.Println(bit)
	
	res := getSum(bit, 5)
	fmt.Println(res)
	
	//updateBIT(bit, len(input), 2, 5)
	//fmt.Println(bit)
	//res = getSum(bit, 5)
	//fmt.Println(res)
	
	q := query(bit, 1,4)
	fmt.Println(q) 
}
