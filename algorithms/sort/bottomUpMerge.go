package main

import (
	"fmt"
)

func main() {
	arr := []int{2,4,5,3,7,9,6,8}
	iterativeMergeSort(arr)
	fmt.Println(arr)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	
	return b
}

func iterativeMergeSort(arr []int) {
	size := len(arr) - 1
	for curSize := 1; curSize <= size; curSize = 2*curSize {
		for leftSize := 0; leftSize < size; leftSize += 2*curSize {
			mid := min(leftSize + curSize - 1, size)
			rightSize := min(leftSize + 2*curSize - 1, size)
			
			merge(arr, leftSize, mid, rightSize) 
		}
	}
}

func merge(arr []int, l, m, r int) {
	var i,j,k int
	n1 := m-l+1
	n2 := r-m
	
	L := make([]int, n1)
	R := make([]int, n2)
	
	for i = 0; i<n1; i++ {
		L[i] = arr[l + i]
	}
	
	for j = 0; j<n2; j++ {
		R[j] = arr[m+1+j]
	}
	
	i = 0
	j = 0
        // Initialize k = low
	k = l
	
	for ; i<n1 && j<n2;  {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		k++
	}
	
	for ; i<n1; i++ {
		arr[k] = L[i]
		k++
	}
	
	for ; j<n2; j++ {
		arr[k] = R[j]
		k++
	}
}
