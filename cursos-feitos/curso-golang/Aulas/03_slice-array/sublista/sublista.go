package main

import "fmt"

func main() {

	listaToda := []int{2, 10, 9, 4, 5, 8, 1, 3, 6, 7}

	var listaMenor5 = make([]int, 0)
	// fmt.Println("Inicial: ", listaMenor5)

	for i := 0; i < len(listaToda); i++ {

		if listaToda[i] <= 5 {
			listaMenor5 = append(listaMenor5, listaToda[i])
			fmt.Println("Posição [", i, "]", listaMenor5)
		}
	}

	// fmt.Println("Lista Toda: ", listaToda)
	// fmt.Println("Lista Menor que Cinco: ", listaMenor5)

}

// {2, 10, 9, 4, 5, 8, 1, 3}
