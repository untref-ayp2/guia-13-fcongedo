package binarysearchtree

import (
	"fmt"
)

type Ordered interface {
	~float64 | ~int | ~string
}

type BinaryNode[T Ordered] struct {
	data  T
	left  *BinaryNode[T]
	right *BinaryNode[T]
}

func NewBinaryNode[T Ordered](data T, left *BinaryNode[T], right *BinaryNode[T]) *BinaryNode[T] {
	return &BinaryNode[T]{left: left, right: right, data: data}
}

func (n *BinaryNode[T]) GetData() T {
	return n.data
}

func (n *BinaryNode[T]) String() string {
	return fmt.Sprintf("%v", n.data)
}
