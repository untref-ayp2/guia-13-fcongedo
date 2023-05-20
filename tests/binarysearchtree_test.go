package tests

import (
	"guia13/binarysearchtree"
	"testing"
)

func TestTamañoVacio(t *testing.T) {
	bstree := binarysearchtree.NewBinarySearchTree[int]()
	if bstree.Size() != 0 {
		t.Errorf("Error el tamaño deberia dar %v, pero dio %v", 0, bstree.Size())
	}
}

func TestTamañoDeSoloRaiz(t *testing.T) {
	bstree := binarysearchtree.NewBinarySearchTree[int]()
	bstree.Insert(4)
	if bstree.Size() != 1 {
		t.Errorf("Error el tamaño deberia dar %v, pero dio %v", 1, bstree.Size())
	}
}

func TestTamañoDeSoloRaizConHijoIzquierdo(t *testing.T) {
	bstree := binarysearchtree.NewBinarySearchTree[int]()
	bstree.Insert(4)
	bstree.Insert(2)
	if bstree.Size() != 2 {
		t.Errorf("Error el tamaño deberia dar %v, pero dio %v", 2, bstree.Size())
	}
}

func TestTamañoDeSoloRaizConHijoDerecho(t *testing.T) {
	bstree := binarysearchtree.NewBinarySearchTree[int]()
	bstree.Insert(4)
	bstree.Insert(7)
	if bstree.Size() != 2 {
		t.Errorf("Error el tamaño deberia dar %v, pero dio %v", 2, bstree.Size())
	}
}

func TestTamañoDeRaizYAmbosHijos(t *testing.T) {
	bstree := binarysearchtree.NewBinarySearchTree[int]()
	bstree.Insert(4)
	bstree.Insert(2)
	bstree.Insert(7)
	if bstree.Size() != 3 {
		t.Errorf("Error el tamaño deberia dar %v, pero dio %v", 3, bstree.Size())
	}
}

func TestRecorridoInOrder(t *testing.T) {
	bstree := binarysearchtree.NewBinarySearchTree[int]()
	bstree.Insert(4)
	bstree.Insert(2)
	bstree.Insert(7)
	if bstree.InOrder() != "2 4 7 " {
		t.Errorf("Error el recorridoInOrder deberia dar 2 4 7, pero dio %v", bstree.InOrder())
	}
}
func TestRemoveRaiz(t *testing.T) {
	bstree := binarysearchtree.NewBinarySearchTree[int]()
	bstree.Insert(4)
	bstree.Insert(2)
	bstree.Insert(7)
	bstree.Remove(4)
	if bstree.InOrder() != "2 7 " {
		t.Errorf("Error el árbol deberia quedar como 2 7, pero dio %v", bstree.String())
	}
}
