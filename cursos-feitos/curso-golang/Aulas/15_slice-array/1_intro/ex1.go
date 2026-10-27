package main

import "fmt"

// copia de slices e arrays de forma segura
func main() {
	var (
		slice = []int{1, 2, 3, 4, 5}
		// any = interface{}
		array = [5]int{1, 2, 3, 4, 5}
	)
	fmt.Println("Tamanho do array:", len(array))
	fmt.Println("Tamanho do array:", cap(array))
	fmt.Println("Tamanho do slice:", len(slice))
	fmt.Println("Tamanho do slice:", cap(slice))

	fmt.Println("==============================")
	slice = append(slice, 6)
	fmt.Println("Tamanho do slice:", len(slice))
	fmt.Println("Tamanho do slice:", cap(slice))

	// appende não funciona com array

	printSlice(slice)

}

func printSlice(slice []int) {
	for _, i := range slice {
		fmt.Println(i)
	}
}

// slice é a fatia de um array
// go cria um array
