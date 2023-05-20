package ejercicios

import (
	"fmt"
	"strings"
)

type BinarySearchTree struct {
	root *BinaryNode
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{root: nil}
}

func (bst *BinarySearchTree) GetRoot() *BinaryNode {
	return bst.root
}

func (bst BinarySearchTree) String() string {
	return bst.InOrder()
}

func (bst BinarySearchTree) InOrder() string {
	sb := strings.Builder{}
	bst.inOrderByNode(&sb, bst.root)
	return sb.String()
}

func (bst BinarySearchTree) inOrderByNode(sb *strings.Builder, root *BinaryNode) {
	if root == nil {
		return
	}
	bst.inOrderByNode(sb, root.left)
	sb.WriteString(fmt.Sprintf("%s ", root))
	bst.inOrderByNode(sb, root.right)
}

func (bst *BinarySearchTree) Insert(k int) {
	bst.root = bst.insertByNode(bst.root, k)
}

func (bst BinarySearchTree) insertByNode(node *BinaryNode, k int) *BinaryNode {
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

func (bst *BinarySearchTree) Search(k int) *BinaryNode {
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

func (bst *BinarySearchTree) FindMin() *BinaryNode {
	if bst.root == nil {
		return nil
	}
	nextLeft := bst.root
	for nextLeft.left != nil {
		nextLeft = nextLeft.left
	}
	return nextLeft
}

func (bst *BinarySearchTree) Remove(k int) {
	bst.removeByNode(bst.root, k)
}

func (bst *BinarySearchTree) removeByNode(root *BinaryNode, k int) *BinaryNode {
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

func (bst *BinarySearchTree) Size() int {
	return size(bst.root)
}

func size(node *BinaryNode) int {
	if node == nil {
		return 0
	}
	return 1 + size(node.left) + size(node.right)
}

func (bst *BinarySearchTree) Sum() int {
	return sum(bst.root)
}

func sum(node *BinaryNode) int {
	if node == nil {
		return 0
	}
	return node.data + sum(node.left) + sum(node.right)
}
