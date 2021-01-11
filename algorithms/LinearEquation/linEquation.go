package main

import (
	"fmt"
)

func main() {
	// 2x + 15y = 17
	// x + 19y = 20
	
	coe := [][]int{{2,15}, {1,19}}
	fmt.Println(coe)
	con := [][]int{{17},{20}}
	fmt.Println(con)
	res := linEquationSolver(coe, con)
	fmt.Println(res)
	
}

// Solve Linear Equation
func linEquationSolver(mat1 [][]int, mat2 [][]int) []int{

        // X = A^-1 * B
	d := fd(mat1)
	in := inMatrix(mat1)
	resMat := mulMatrix(in, mat2)
	res := divMatrix(resMat, d)
	return res 
}

// Calculate Determinaant
func fd(in [][]int) int {
	l := len(in)
	fir := 1
	sec := 1
	for i := 0; i < l; i++{
		for j := 0; j < l; j++{
			if i == j {
				fir *= in[i][j]
			} else {
				sec *= in[i][j]
			} 
		}
	}
	
	return fir - sec
}

// Multiply 2 matrix
func mulMatrix (mat1 [][]int, mat2 [][]int) []int {
	l := len(mat1)
	fir := 0
	sec := 0
	res := 0
	out := make([]int,2)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if (i == j) {
				fir = mat1[i][j] * mat2[j][0] 
			} else {
				sec = mat1[i][j] * mat2[j][0]
			}
		}
		res = fir + sec
		out[i] = res
		fir = 0
		sec = 0 	
	}
	return out
}

// Calculate inverse of matrix
func inMatrix(mat [][]int) [][]int {
	l := len(mat)
	
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if (i == j) {
				if (j == l-1) {
					break
				}
				mat[i][j], mat[i+1][j+1] = mat[i+1][j+1], mat[i][j]
			} else {
				mat[i][j] = -mat[i][j]
			}
		}
	}
	
	return mat
}

// Divide the matrix 
func divMatrix(mat []int, d int) []int {
	l := len(mat)
	
	for i := 0; i < l; i++ {
		mat[i] = mat[i]/d
	}
	
	return mat
}
