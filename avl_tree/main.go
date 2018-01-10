package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

// node
type node struct {
	id          int
	height      int
	left, right *node
}

func (n *node) String() string {
	return fmt.Sprintf("%d", n.id)
}

func height(n *node) int {
	if n == nil {
		return 0
	}

	return n.height
}
func (n *node) balance() int {
	if n == nil {
		return 0
	}

	return height(n.left) - height(n.right)
}

// queue
type queue struct {
	l []*node
}

func (q *queue) pop_front() (n *node) {
	n = q.l[0]
	q.l = q.l[1:]
	return
}

func (q *queue) push_back(n *node) {
	q.l = append(q.l, n)
}

func (q queue) size() int {
	return len(q.l)
}

func minDepth(root *node) int {
	if root == nil {
		return 0
	}

	if root.left == nil && root.right == nil {
		return 1
	}

	if root.left == nil {
		return minDepth(root.right)
	}

	if root.right == nil {
		return minDepth(root.left)
	}

	l := minDepth(root.left) + 1
	r := minDepth(root.right) + 1
	if l < r {
		return l
	}
	return r
}

func size(root *node) int {
	if root == nil {
		return 0
	}

	return 1 + size(root.left) + size(root.right)
}

func preorderTraversal(n *node) {
	if n != nil {
		log.Print(n.id)
		preorderTraversal(n.left)
		preorderTraversal(n.right)
	}
}

func dfTraversal(n *node) {
	if n == nil {
		return
	}

	log.Print(n.id)
	dfTraversal(n.left)
	dfTraversal(n.right)
}

func bfTraversal(root *node) {
	if root == nil {
		return
	}

	q := queue{}
	q.push_back(root)

	for q.size() > 0 {
		n := q.pop_front()
		if n != nil {
			log.Print(n.id)
			q.push_back(n.left)
			q.push_back(n.right)
		}
	}
}

func power(n int) (res int) {
	res = 1
	for i := 1; i < n; i++ {
		res *= 2
	}

	return
}

func inStructured(r io.Reader) *node {
	noderows := make([][]*node, 0)
	idx := 0

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		idx++

		fields := strings.Fields(scanner.Text())
		if len(fields) != power(idx) {
			log.Fatalf("invalid input: %d %v", len(fields), fields)
		}

		noderow := make([]*node, 0)
		for _, v := range fields {
			n, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalf("invalid input: %v", fields)
			}

			if n == -1 {
				noderow = append(noderow, nil)
			} else {
				noderow = append(noderow, &node{id: n, height: 1})
			}
		}

		noderows = append(noderows, noderow)
	}

	for i := 0; i < len(noderows)-1; i++ {
		for j := 0; j < len(noderows[i]); j++ {
			if noderows[i][j] != nil {
				noderows[i][j].left = noderows[i+1][j*2]
				noderows[i][j].right = noderows[i+1][j*2+1]
			}
		}
	}

	if len(noderows) > 0 {
		return noderows[0][0]
	}
	return nil
}

func max(l, r int) int {
	if l > r {
		return l
	}

	return r
}

func rightRotate(n *node) (res *node) {
	res = n.left
	tmp := res.right

	res.right = n
	n.left = tmp

	n.height = max(height(n.left), height(n.right)) + 1
	res.height = max(height(res.left), height(res.right)) + 1

	return
}

func leftRotate(n *node) (res *node) {
	res = n.right
	tmp := res.left

	res.left = n
	n.right = tmp

	n.height = max(height(n.left), height(n.right)) + 1
	res.height = max(height(res.left), height(res.right)) + 1

	return
}

func insert(root *node, key int) *node {
	if root == nil {
		return &node{id: key, height: 1}
	}

	if key < root.id {
		root.left = insert(root.left, key)
	} else if key > root.id {
		root.right = insert(root.right, key)
	} else {
		return root
	}

	root.height = 1 + max(height(root.left), height(root.right))

	b := root.balance()

	if b > 1 && key < root.left.id {
		return rightRotate(root)
	}

	if b < -1 && key > root.right.id {
		return leftRotate(root)
	}

	if b > 1 && key > root.left.id {
		root.left = leftRotate(root.left)
		return rightRotate(root)
	}

	if b < -1 && key < root.right.id {
		root.right = rightRotate(root.right)
		return leftRotate(root)
	}

	return root
}

func inInsert(in string) (root *node) {
	for _, v := range strings.Fields(in) {
		n, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("invalid input: %v", v)
		}

		root = insert(root, n)
	}

	return
}

func main() {
	{
		tree := inInsert("1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31")
		log.Print("preorder of tree")
		preorderTraversal(tree)
	}

	{
		tree := inInsert("12 8 18 5 11 17 4 7")
		log.Print("preorder of tree")
		preorderTraversal(tree)
	}

	{
		bb := bytes.NewBufferString(`1
2 3
4 5 6 7
`)
		tree := inStructured(bb)
		//tree := inStructured(os.Stdin)

		log.Print("depth first traversal")
		dfTraversal(tree)

		log.Print("breadth first traversal")
		bfTraversal(tree)

		log.Printf("min depth of tree: %d", minDepth(tree))
		log.Printf("size of tree: %d", size(tree))

		log.Print("preorder of tree")
		preorderTraversal(tree)
	}
}
