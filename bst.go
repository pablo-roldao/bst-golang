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

func (tree *Node) ToPreOrderString() string {
	var result string
	if tree != nil {
		result += fmt.Sprintf("[%d]", tree.value)
		result += tree.left.ToPreOrderString()
		result += tree.right.ToPreOrderString()
	}

	return result
}
