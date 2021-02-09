package main

import (
	"fmt"
)

const size int = 100
var tree [2*size]int

func build(arr []int) {
	n := len(arr)
	
	for i := 0; i<n; i++ {
		tree[n + i] = arr[i]
	}
	
	for i := n-1; i > 0; i-- {
		tree[i] = tree[2*i] + tree[2*i+1]
	}
}

func query(arr []int, l, h int) int {
	res := 0
	n := len(arr)
	
	l = l+n
	h = h+n
	
	for ; l<h; {
		// golang if must have bool condition
		// if 0 is not valid in go
		if l&1 != 0 {
			res = res + tree[l]
			l = l+1 
		}
		
		// golang if must have bool condition
		// if 0 is not valid in go
		if h&1 != 0 {
			h = h-1
			res = res + tree[h]
		}
		
		// divide l and h by 2
		// l >>=1 ==> l = l/2
		// h >>=1 ==> h = h/2
		l >>= 1
		h >>= 1

	}
	
	return res
}


func updateNode(arr []int, index, k int) {
	n := len(arr)
	
	tree[index + n] = k
	index = index + n
	
	for i := index; i>1; {
		// i = i>>1 ==> i = i/2
		// i = i^1 == i+1 ==> if i is even
		// i = i^1 == i-1 ==> if i is odd
		tree[i>>1] = tree[i] + tree[i^1]
		i = i>>1
	}	
}

func main() {
	arr := []int{1,2,3,4,5,6,7,8,9,10,11,12}
	build(arr)
	fmt.Println(tree)
	resp := query(arr,1,3)
	fmt.Println(resp)
	updateNode(arr, 2, 19)
	fmt.Println(tree)
	resp = query(arr,1,3)
	fmt.Println(resp)
}
