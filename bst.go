package bst

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func (tree *Node) Insert(value int) *Node {
	if tree == nil {
		return &Node{value, nil, nil}
	}

	if value < tree.value {
		tree.left = tree.left.Insert(value)
	} else if value > tree.value {
		tree.right = tree.right.Insert(value)
	}

	return tree
}

func (tree *Node) PrintPreOrder() {
	if tree != nil {
		fmt.Printf("[%d]", tree.value)
		tree.left.PrintPreOrder()
		tree.right.PrintPreOrder()
	}
}
