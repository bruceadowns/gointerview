package main

import "fmt"

type node struct {
	id          int
	left, right *node
}

func (n *node) String() string {
	return fmt.Sprintf("%d", n.id)
}

func size(root *node) int {
	if root == nil {
		return 0
	}

	return size(root.left) + 1 + size(root.right)
}

func main() {
	tree6 := &node{6, nil, nil}
	tree7 := &node{7, nil, nil}
	tree4 := &node{4, nil, nil}
	tree5 := &node{5, tree6, tree7}
	tree2 := &node{2, tree4, tree5}
	tree3 := &node{3, nil, nil}
	tree1 := &node{1, tree2, tree3}

	fmt.Println(size(tree1))
}
