package main

import (
	"fmt"
)

type node struct {
	value int
	balance int
	left  *node
	right *node
}

func main() {
	var root *node
	root = insert(root, 56)
	fmt.Println(root)
	root = insert(root, 34)
	
	root = insert(root, 20)
	fmt.Println(root)
	inorder(root) 
}


func balcalculator() func(*node) int {
	left := 0
	right := 0
	balance := 0
	var calculatebal func(root *node) int
	calculatebal = func(root *node) int {
				var LB func(*node)
				var RB func(*node) 
	
				LB = func(root *node) {
					fmt.Println("root is:", root)
					if root.left != nil {
						left += 1
						LB(root.left)
					}
				}
	
				RB = func(root *node) {
					if root.right != nil {
						right += 1
						RB(root.right)
					}
				}
				LB(root)
				RB(root)
				balance = left - right
				return balance
			}
			
	return calculatebal
}

func insert(root *node, val int) *node {
	if root == nil {
		bst := &node{value: val, balance: 0, left: nil, right: nil}
		return bst
	}
	bc := balcalculator()
	if val <= root.value {
		if root.left == nil {
			root.left = &node{value: val, left:nil, right:nil}
			root.left.balance = bc(root.left)
		}else{
			root.left = insert(root.left, val)
		}
	}
	root.balance = bc(root)
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
	fmt.Println(root.value, root.balance)
	if root.right != nil {
		inorder(root.right)
	}
}
