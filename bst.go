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

func (tree *Node) ToInOrderString() string {
	var result string
	if tree != nil {
		result += tree.left.ToInOrderString()
		result += fmt.Sprintf("[%d]", tree.value)
		result += tree.right.ToInOrderString()
	}

	return result
}

func (tree *Node) ToPosOrderString() string {
	var result string
	if tree != nil {
		result += tree.left.ToPosOrderString()
		result += tree.right.ToPosOrderString()
		result += fmt.Sprintf("[%d]", tree.value)
	}

	return result
}

func (tree *Node) ToReverseString() string {
	var result string
	if tree != nil {
		result += tree.right.ToReverseString()
		result += fmt.Sprintf("[%d]", tree.value)
		result += tree.left.ToReverseString()
	}

	return result
}

func (tree *Node) GetHeight() int {
	if tree == nil {
		return 0
	}

	leftHeight := tree.left.GetHeight()
	rightHeight := tree.right.GetHeight()

	if leftHeight > rightHeight {
		return leftHeight + 1
	} else {
		return rightHeight + 1
	}
}

func (tree *Node) GetNumberOfLeaves() int {
	if tree == nil {
		return 0
	}

	result := 0
	result += tree.left.GetNumberOfLeaves()
	result += tree.left.GetNumberOfLeaves()

	if tree.left == nil && tree.right == nil {
		result++
	}

	return result
}
