package main

import (
	"fmt"

	"guia13/binarysearchtree"
)

func main() {

	fmt.Println("Binary Search Tree")
	bst := binarysearchtree.NewBinarySearchTree[int]()
	bst.Insert(7)
	bst.Insert(5)
	bst.Insert(9)
	bst.Insert(2)
	bst.Insert(6)
	bst.Insert(8)
	bst.Insert(12)
	bst.Insert(3)
	bst.Insert(10)
	bst.Insert(10)
	bst.Insert(10)
	bst.Insert(10)
	fmt.Println(bst)

	node := bst.GetRoot()
	fmt.Println(node.GetData())

	fmt.Println(bst.InOrder())

	fmt.Println("Min", bst.FindMin().GetData())

	fmt.Println(bst.Search(3))
	fmt.Println(bst.Search(2))
	fmt.Println(bst.Search(5))
	fmt.Println(bst.Search(12))
	fmt.Println(bst.Search(10))
	fmt.Println(bst.Search(4))
	fmt.Println("Completo")
	fmt.Println("Tamaño")
	fmt.Println(bst.Size())
	bst.Remove(7)
	fmt.Println(bst)
	fmt.Println(bst.GetRoot().GetData())

}
