package main

import "fmt"

type node struct {
	data        int
	left, right *node
}

func lca(root, n1, n2 *node) *node {
	if root == nil {
		return nil
	}

	if root.data == n1.data || root.data == n2.data {
		return root
	}

	ll := lca(root.left, n1, n2)
	lr := lca(root.right, n1, n2)
	if ll != nil && lr != nil {
		return root
	}
	if ll != nil {
		return ll
	}
	return lr
}

func level(root, n *node, l int) int {
	if root == nil {
		return -1
	}

	if root.data == n.data {
		return l
	}

	levelL := level(root.left, n, l+1)
	if levelL != -1 {
		return levelL
	}

	levelR := level(root.right, n, l+1)
	if levelR != -1 {
		return levelR
	}

	return -1
}

func main() {
	six := &node{6, nil, nil}
	seven := &node{7, nil, nil}
	four := &node{4, nil, nil}
	five := &node{5, six, seven}
	two := &node{2, four, five}
	three := &node{3, nil, nil}
	root := &node{1, two, three}

	firstNode := three
	secondNode := seven

	firstLevel := level(root, firstNode, 0)
	secondLevel := level(root, secondNode, 0)

	l := level(root, lca(root, firstNode, secondNode), 0)
	dist := firstLevel + secondLevel - 2*l
	fmt.Printf("%d", dist)
}
