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

// In-Order 1) hijo izquierdo 2) raiz 3) hijo derecho

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

// Pre-Order 1) raiz 2) hijo izquierdo 3) hijo derecho
func (bst BinarySearchTree[T]) PreOrder() string {
	sb := strings.Builder{}
	bst.preOrderByNode(&sb, bst.root)
	return sb.String()
}

func (bst BinarySearchTree[T]) preOrderByNode(sb *strings.Builder, root *BinaryNode[T]) {
	if root == nil {
		return
	}
	sb.WriteString(fmt.Sprintf("%s ", root))
	bst.preOrderByNode(sb, root.left)
	bst.preOrderByNode(sb, root.right)
}

// Post-Order 1) hijo izquierdo 2) hijo derecho 3) raiz
func (bst BinarySearchTree[T]) PostOrder() string {
	sb := strings.Builder{}
	bst.postOrderByNode(&sb, bst.root)
	return sb.String()
}

func (bst BinarySearchTree[T]) postOrderByNode(sb *strings.Builder, root *BinaryNode[T]) {
	if root == nil {
		return
	}
	bst.postOrderByNode(sb, root.left)
	bst.postOrderByNode(sb, root.right)
	sb.WriteString(fmt.Sprintf("%s ", root))
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

//Escribir un método recursivo que calcule (y devuelva) la altura de un árbol binario de búsqueda.
// El método debe devolver -1 si el árbol está vacío, y 0 si solo contiene la raiz

func (bst *BinarySearchTree[T]) Height() int {
	return bst.height(bst.root) - 1
}

func (bst *BinarySearchTree[T]) height(node *BinaryNode[T]) int {

	if node == nil {
		return 0

	}
	altura := max(bst.height(node.left), bst.height(node.right)+1)

	return altura
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

 func (bst *BinarySearchTree[T]) SumarHojas() int {
 	if bst.root == nil {
 		return 0
 	}
 	return bst.sumarHojas(bst.root)
 }

// func (bst *BinarySearchTree[T]) sumaHojas(node *BinaryNode[T]) int {
// 	if node.left == nil && node.right == nil {
// 		return node.data
// 	}
// 	suma := 0
// 	if node.left != nil {
// 		suma += bst.sumaHojas(node.left)
// 	}
// 	if node.right != nil {
// 		suma += bst.sumaHojas(node.right)
// 	}
// 	return suma
// }

//Escribir un método recursivo que calcule (y devuelva) la cantidad de nodos que son hojas de un árbol binario de búsqueda.

func (bst *BinarySearchTree[T]) ContarNodosQueSonHojas() int {
	return bst.contarNodosQueSonHojas(bst.root)
}

func (bst *BinarySearchTree[T]) contarNodosQueSonHojas(node *BinaryNode[T]) int {
	if node == nil {
		return 0
	}
	if node.left == nil && node.right == nil {
		// El nodo es una hoja
		return 1
	}
	nodosHojasIzquierda := bst.contarNodosQueSonHojas(node.left)
	nodosHojasDerecha := bst.contarNodosQueSonHojas(node.right)
	return nodosHojasIzquierda + nodosHojasDerecha
}

//Escribir un método recursivo que calcule (y devuelva) la cantidad de nodos internos(que no son hojas ni raiz) de un árbol binario de búsqueda.

func (bst *BinarySearchTree[T]) ContarNodosInternos() int {
	if bst.root == nil {
		return 0
	}
	return bst.contarNodosInternos(bst.root)
}

func (bst *BinarySearchTree[T]) contarNodosInternos(node *BinaryNode[T]) int {
	if node == nil || (node.left == nil && node.right == nil) {
		return 0
	}
	return 1 + bst.contarNodosInternos(node.left) + bst.contarNodosInternos(node.right)
}

// func (bst *BinarySearchTree) SumaParImpar() (int, int) {
// 	p, i := sumaParImpar(bst.root)
// 	return p, i
// }
// func sumaParImpar(node *BinaryNode) (int, int) {
// 	if node == nil {
// 		return 0, 0
// 	}
// 	pR, iR := 0, 0
// 	if node.data%2 == 0 {
// 		pR = node.data
// 	} else {
// 		iR = node.data
// 	}
// 	pI, iI := sumaParImpar(node.left)
// 	pD, iD := sumaParImpar(node.right)
// 	return (pR + pI + pD), (iR + iI + iD)
// }
//Escribir un método recursivo que calcule (y devuelva) la cantidad de nodos von valores pares e impares de un árbol binario de búsqueda.

// func (bst *BinarySearchTree[T]) ContarNodosParesImpares() (int, int) {
// 	return bst.contarNodosParesImpares(bst.root)
// }

// func (bst *BinarySearchTree[T]) contarNodosParesImpares(node *BinaryNode[T]) (int, int) {
// 	if node == nil {
// 		return 0, 0
// 	}

// 	paresIzquierda, imparesIzquierda := bst.contarNodosParesImpares(node.left)
// 	paresDerecha, imparesDerecha := bst.contarNodosParesImpares(node.right)

// 	pares := paresIzquierda + paresDerecha
// 	impares := imparesIzquierda + imparesDerecha

// 	if node.data%2 == 0 {
// 		pares++
// 	} else {
// 		impares++
// 	}

// 	return pares, impares
// }
