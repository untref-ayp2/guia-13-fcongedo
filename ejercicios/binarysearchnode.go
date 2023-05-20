package ejercicios

import (
	"fmt"
)

type BinaryNode struct {
	data  int
	left  *BinaryNode
	right *BinaryNode
}

func NewBinaryNode(data int, left *BinaryNode, right *BinaryNode) *BinaryNode {
	return &BinaryNode{left: left, right: right, data: data}
}

func (n *BinaryNode) GetData() int {
	return n.data
}

func (n *BinaryNode) String() string {
	return fmt.Sprintf("%v", n.data)
}
