package structures

type BinaryTreeNode struct {
	Value int
	Left *BinaryTreeNode
	Right *BinaryTreeNode
}

func NewBinaryTreeNode(value int) *BinaryTreeNode {
	return &BinaryTreeNode{value, nil, nil}
}