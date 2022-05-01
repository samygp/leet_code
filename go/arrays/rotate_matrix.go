package arrays

import (
	"fmt"
)

func RotateMatrix() {
	var daMatrix [][]int32 = [][]int32{
		{1, 2, 3, 4, 5},
		{16, 17, 18, 19, 6},
		{15, 24, 25, 20, 7},
		{14, 23, 22, 21, 8},
		{13, 12, 11, 10, 9},
	}

	printMatrix(daMatrix)
	layers := len(daMatrix) / 2
	//layers := int(math.Floor()
	for l := 0; l < layers; l++ {
		rotateLeft(l, daMatrix)
	}
	printMatrix(daMatrix)

}

func rotateLeft(layer int, daMatrix [][]int32) {
	n := len(daMatrix) - 1 - layer
	size := n - layer

	for i := 0; i < size; i++ {
		first := daMatrix[layer][layer+i]
		// arriba
		daMatrix[layer][layer+i] = daMatrix[layer+i][n]
		// derecha
		daMatrix[layer+i][n] = daMatrix[n][n-i]
		// abajo
		daMatrix[n][n-i] = daMatrix[n-i][layer]
		// izquierda
		daMatrix[n-i][layer] = first
	}
}
func printMatrix(matrix [][]int32) {
	fmt.Printf("\n")
	for _, line := range matrix {
		fmt.Println(line)
	}
}
