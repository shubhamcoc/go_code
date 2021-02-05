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
	nodeValue := []int{5,6,3,19,11,25,17,2,4}
	for _, ele := range nodeValue {
		root = insert(root, ele)
	}
	inorder(root)
	num := 19
	found := find(root, num)
	if !found {
		fmt.Println("key not found :", num)
	}
}

// function to create and insert element in Binary tree
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

// function to do inorder traversal in Binary tree
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

// function to search element in binary tree
func find(root *node, key int) bool {
	if root == nil {
		fmt.Println("Binary tree is empty")
	}
	
	if root.value == key {
		fmt.Println("Key found:", key)
		return true
	}
	var res bool
	if key < root.value {
		if root.left != nil {
			res = find(root.left, key)
		}else{
			res = false
		}
	}else{
		if root.right != nil {
			res = find(root.right, key)
		}else{
			res = false
		}
	}
	
	return res
}
