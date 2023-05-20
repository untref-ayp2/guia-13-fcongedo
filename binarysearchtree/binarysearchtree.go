package binarysearchtree

import (
	"fmt"
	"strings"
)

type BinarySearchTree[T Ordered] struct {
	root *BinaryNode[T]
}

func NewBinarySearchTree[T Ordered]() *BinarySearchTree[T] {
	return &BinarySearchTree[T]{root: nil}
}

func (bst *BinarySearchTree[T]) GetRoot() *BinaryNode[T] {
	return bst.root
}

func (bst BinarySearchTree[T]) String() string {
	return bst.InOrder()
}

func (bst BinarySearchTree[T]) InOrder() string {
	sb := strings.Builder{}
	bst.inOrderByNode(&sb, bst.root)
	return sb.String()
}

func (bst BinarySearchTree[T]) inOrderByNode(sb *strings.Builder, root *BinaryNode[T]) {
	if root == nil {
		return
	}
	bst.inOrderByNode(sb, root.left)
	sb.WriteString(fmt.Sprintf("%s ", root))
	bst.inOrderByNode(sb, root.right)
}

func (bst *BinarySearchTree[T]) Insert(k T) {
	bst.root = bst.insertByNode(bst.root, k)
}

func (bst BinarySearchTree[T]) insertByNode(node *BinaryNode[T], k T) *BinaryNode[T] {
	if node == nil {
		return NewBinaryNode(k, nil, nil)
	}
	if k < node.data {
		node.left = bst.insertByNode(node.left, k)
	} else if k > node.data {
		node.right = bst.insertByNode(node.right, k)
	}
	return node
}

func (bst *BinarySearchTree[T]) Search(k T) *BinaryNode[T] {
	encontrado := false
	node := bst.root
	for node != nil && !encontrado {
		if k < node.data {
			node = node.left
		} else if k > node.data {
			node = node.right
		} else {
			encontrado = true
			return node
		}
	}
	return nil
}

func (bst *BinarySearchTree[T]) FindMin() *BinaryNode[T] {
	if bst.root == nil {
		return nil
	}
	nextLeft := bst.root
	for nextLeft.left != nil {
		nextLeft = nextLeft.left
	}
	return nextLeft
}

func (bst *BinarySearchTree[T]) Remove(k T) {
	bst.removeByNode(bst.root, k)
}

func (bst *BinarySearchTree[T]) removeByNode(root *BinaryNode[T], k T) *BinaryNode[T] {
	if root == nil {
		return root
	}
	if k > root.data {
		root.right = bst.removeByNode(root.right, k)
	} else if k < root.data {
		root.left = bst.removeByNode(root.left, k)
	} else {
		if root.left == nil {
			return root.right
		} else {
			temp := root.left
			for temp.right != nil {
				temp = temp.right
			}
			root.data = temp.data
			root.left = bst.removeByNode(root.left, temp.data)
		}
	}
	return root
}

func (bst *BinarySearchTree[T]) Size() int {
	return size(bst.root)
}

func size[T Ordered](node *BinaryNode[T]) int {
	if node == nil {
		return 0
	}
	return 1 + size(node.left) + size(node.right)
}
