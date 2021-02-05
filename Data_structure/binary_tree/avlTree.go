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
	root = insert(root, 37)
	
	//root = insert(root, 20)
	//root = insert(root, 60)
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
					//fmt.Println("root is:", root)
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
	}else {
		if root.right == nil {
			root.right = &node{value: val, left:nil, right:nil}
			root.right.balance = bc(root.right)
		}else{
			root.right = insert(root.right, val)
		} 
	}
	root.balance = bc(root)
	root1 := balanceTree(root)
	// balance is incorrect in case of new root
	if root1 != root {
		fmt.Println("root1 is:", root1)
		root1.balance = bc(root1)
		root1.left.balance = bc(root1.left)
		root1.right.balance = bc(root1.right)
		return root1
	}
	return root
}

func balanceTree(root *node) *node {
	switch root.balance {
		case 0:
			break
		case 1:
			break
		case -1:
			break
		default:
			root = balancer(root)
	}
	return root
}

func balancer(root *node) *node {
	if root.left != nil {
		if root.left.left != nil {
			temp := root
			root = root.left
			root.right = temp 
		}
		
		if root.left.right != nil {
			temp1 := root.left
			temp2 := root
			root = root.left.right
			root.left = temp1
			root.right = temp2
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
	fmt.Println(root.value, root.balance)
	if root.right != nil {
		inorder(root.right)
	}
}
