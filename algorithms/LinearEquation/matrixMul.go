package main

import (
	"fmt"
	"sync"
)

func main() {
	A := [][]int{
		{1,2,3,4,5},
		{2,3,4,5,6},
		{3,4,5,6,7},
		{4,5,6,7,8},
		{5,6,7,8,9},
	}
	B := [][]int{
		{1,2,3,4,5},
		{2,3,4,5,6},
		{3,4,5,6,7},
		{4,5,6,7,8},
		{5,6,7,8,9},
	}
	fmt.Println(len(A))
	res := matrixMultiplication(A, B)
	fmt.Println(res)
	res1 := Multiply(A, B)
	fmt.Println(res1)
}

func matrixMultiplication(A, B [][]int) []int {
	size := len(A)
	C := make([]int,5)
	var wg sync.WaitGroup
	f := func(A, B [][]int, i int, size int) {	
		j := i
		C[i] = 0
			
		for k := 0; k<size; k++ {
			C[i] = C[i] + A[i][k] * B[k][j]		
		}
		wg.Done()
	}
	
	for i := 0; i<size; i++ {
		wg.Add(1)
		go f (A, B, i, size)
	}
	wg.Wait()
	return C
}

func Multiply(A, B [][]int) []int {
	size := len(A)
	C := make([]int,5)
	for i := 0; i<size; i++ {
		for j := 0; j<size; j++ {
			//C[i] = make([]int,5)
			C[i] = 0
			for k := 0; k<size; k++ {
				C[i] = C[i] + A[i][k] * B[i][k]
			}	
		}
	}
	return C
}
