package tests

import (
	"fmt"
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

func TestRecorridoPreOrder(t *testing.T) {
	bstree := binarysearchtree.NewBinarySearchTree[int]()
	bstree.Insert(4)
	bstree.Insert(2)
	bstree.Insert(7)
	if bstree.PreOrder() != "4 2 7 " {
		t.Errorf("Error el recorridoInOrder deberia dar 4 2 7, pero dio %v", bstree.PreOrder())
	}
}

func TestRecorridoPostOrder(t *testing.T) {
	bstree := binarysearchtree.NewBinarySearchTree[int]()
	bstree.Insert(4)
	bstree.Insert(2)
	bstree.Insert(7)
	if bstree.PostOrder() != "2 7 4 " {
		t.Errorf("Error el recorridoInOrder deberia dar 2 7 4, pero dio %v", bstree.PostOrder())
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

func TestBinarySearchTree_Height(t *testing.T) {
	// Crea un BinarySearchTree de prueba
	bst := binarysearchtree.NewBinarySearchTree[int]()

	// Inserta los nodos en el árbol
	bst.Insert(10)
	bst.Insert(6)
	bst.Insert(14)
	bst.Insert(4)
	bst.Insert(8)
	bst.Insert(12)
	bst.Insert(16)

	// Calcula la altura esperada del árbol
	expectedHeight := 2

	// Obtiene la altura del árbol utilizando la función Height
	actualHeight := bst.Height()

	// Compara la altura obtenida con la altura esperada
	if actualHeight != expectedHeight {
		t.Errorf("Altura incorrecta. Se esperaba %d, se obtuvo %d", expectedHeight, actualHeight)
	}
}

func TestContarNodosQueSonHojas(t *testing.T) {
	bst := binarysearchtree.NewBinarySearchTree[int]()
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(15)
	bst.Insert(3)
	bst.Insert(7)
	bst.Insert(12)
	bst.Insert(17)

	// El árbol debería ser el siguiente:
	//       10
	//      /  \
	//     5    15
	//    / \   / \
	//   3   7 12  17

	expectedCount := 4 // Hay 4 nodos que son hojas (3, 7, 12, 17)
	actualCount := bst.ContarNodosQueSonHojas()

	if actualCount != expectedCount {
		t.Errorf("El resultado obtenido (%d) no coincide con el resultado esperado (%d)", actualCount, expectedCount)
	}

}

func TestContarNodosInternos(t *testing.T) {
	// Crear el árbol de prueba
	bst := binarysearchtree.NewBinarySearchTree[int]()
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(15)
	bst.Insert(3)
	bst.Insert(7)
	bst.Insert(12)
	bst.Insert(17)

	fmt.Println(bst)

	// Obtener el número de nodos internos
	nodosInternos := bst.ContarNodosInternos()

	// Verificar el resultado esperado
	// El árbol tiene 3 nodos internos: 10, 5, 15
	// Los nodos 3, 7 12 y 17 son hojas y no se cuentan como nodos internos.
	expected := 3
	if nodosInternos != expected {
		t.Errorf("Número incorrecto de nodos internos. Se esperaban %d pero se obtuvo %d", expected, nodosInternos)
	}
}

// func TestCountNodesWithParity(t *testing.T) {
// 	bst := binarysearchtree.NewBinarySearchTree[int]()
// 	bst.Insert(5)
// 	bst.Insert(3)
// 	bst.Insert(7)
// 	bst.Insert(2)
// 	bst.Insert(4)
// 	bst.Insert(6)
// 	bst.Insert(8)

// 	pares, impares := bst.SumaParImpar()
// 	expectedPares := 4
// 	expectedImpares := 3

// 	if pares != expectedPares {
// 		t.Errorf("Cantidad de nodos pares incorrecta. Se esperaba %d, se obtuvo %d", expectedPares, pares)
// 	}

// 	if impares != expectedImpares {
// 		t.Errorf("Cantidad de nodos impares incorrecta. Se esperaba %d, se obtuvo %d", expectedImpares, impares)
// 	}
// }
