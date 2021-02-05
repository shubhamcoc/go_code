package main

import (
	"fmt"
)

type node struct {
	value int
	left  *node
	right *node
}

func main() {
	var root *node
	fmt.Println(root)
	root = insert(root, 5)
	fmt.Println(root)
	root = insert(root, 6)
	fmt.Println(root)
	root = insert(root, 4)
	fmt.Println(root)
	inorder(root)
}

func insert(root *node, val int) *node {
	if root == nil {
		bst := &node{value: val, left: nil, right: nil}
		return bst
	}
	if val <= root.value{
		if root.left == nil {
			root.left = &node{value: val, left: nil, right: nil}
		}else{
			root.left = insert(root.left, val)
		}
	}else{
		if root.right == nil {
			root.right = &node{value: val, left:nil, right: nil}
		}else{
			root.right = insert(root.right, val)
		}
	}
	return root
}

func inorder(root *node) {
	if root == nil {
		fmt.Println("Binary tree is empty")
		return
	}
	
	if root.left != nil {
		inorder(root.left)
	}
	fmt.Println(root.value)
	if root.right != nil {
		inorder(root.right)
	} 
}
