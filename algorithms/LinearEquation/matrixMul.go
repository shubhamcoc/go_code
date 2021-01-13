package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	A := [][]int{
		{1,2,3,4,5,6},
		{2,3,4,5,6,7},
		{3,4,5,6,7,8},
		{4,5,6,7,8,9},
		{5,6,7,8,9,10},
		{6,7,8,9,10,11},
	}
	B := [][]int{
		{1,2,3,4,5,6},
		{2,3,4,5,6,7},
		{3,4,5,6,7,8},
		{4,5,6,7,8,9},
		{5,6,7,8,9,10},
		{6,7,8,9,10,11},
	}
	fmt.Println(len(A))
	t1 := time.Now().Nanosecond()
	res := matrixMultiplication(A, B)
	t2 := time.Now().Nanosecond()
	fmt.Println(res)
	fmt.Println(t2-t1)
	t3 := time.Now().Nanosecond()
	res1 := Multiply(A, B)
	t4 := time.Now().Nanosecond()
	fmt.Println(res1)
	fmt.Println(t4-t3)
}

func matrixMultiplication(A, B [][]int) []int {
	size := len(A)
	C := make([]int,size)
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
		go f(A, B, i, size)
	}
	wg.Wait()
	return C
}

func Multiply(A, B [][]int) []int {
	size := len(A)
	C := make([]int,size)
	for i := 0; i<size; i++ {
		for j := 0; j<size; j++ {
			C[i] = 0
			for k := 0; k<size; k++ {
				C[i] = C[i] + A[i][k] * B[i][k]
			}	
		}
	}
	return C
}
