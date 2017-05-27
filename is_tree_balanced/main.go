package main

import "fmt"

type node struct {
	data        int
	left, right *node
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func (n *node) height() int {
	if n == nil {
		return 0
	}

	return 1 + max(n.left.height(), n.right.height())
}

func isBalanced(n *node) bool {
	if n == nil {
		return true
	}

	lh := n.left.height()
	rh := n.right.height()

	if lh-rh <= 1 && isBalanced(n.left) && isBalanced(n.right) {
		return true
	}

	return false
}

func newNode(i int) *node {
	//n := &node{}
	//n.data = i
	//return n

	return &node{i, nil, nil}
}

func main() {
	/*
		root := newNode(1)
		root.left = newNode(2)
		root.right = newNode(3)
		root.left.left = newNode(4)
		root.left.right = newNode(5)
		root.right.right = newNode(6)
		root.left.right.left = newNode(7)
	*/

	root := newNode(1)
	root.left = newNode(2)
	root.left.left = newNode(4)
	root.left.right = newNode(5)
	root.left.left.left = newNode(8)
	root.left.left.right = newNode(9)
	root.left.right.right = newNode(10)
	root.left.right.right.left = newNode(12)
	root.left.left.right = newNode(9)
	root.left.left.right.left = newNode(11)
	root.right = newNode(3)
	root.right.left = newNode(6)
	root.right.right = newNode(7)

	if isBalanced(root) {
		fmt.Println("balanced")
	} else {
		fmt.Println("not balanced")
	}
}
