package main

import "fmt"

// deep copy
func main() {
	matriz1 := [][]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println("---original---")
	matriz2 := make([][]int, len(matriz1))
	fmt.Println(matriz1)
	fmt.Println(matriz2)

	fmt.Println("--copy errado/Shalow copy---")
	copy(matriz2, matriz1)
	fmt.Println(matriz1)
	fmt.Println(matriz2)

	// matriz2[0][0] = 99
	//
	// fmt.Println(matriz1)
	// fmt.Println(matriz2)
	fmt.Println("---correção/Deep Copy---")

	// // matriz2[0] = []int{99, 99}
	// fmt.Println(matriz1)
	// fmt.Println(matriz2)

	matriz3 := [][]int{
		{1, 2},
		{3, 4},
	}

	matriz4 := deepCopy(matriz3)
	matriz3[0][0] = 99
	fmt.Println("Matriz 3:", matriz3)
	fmt.Println("Matriz 4:", matriz4)

}

func deepCopy(matriz [][]int) [][]int {
	destino := make([][]int, len(matriz))

	for i, slice := range matriz {
		destino[i] = make([]int, len(slice))
		copy(destino[i], slice)

	}
	return destino
}
