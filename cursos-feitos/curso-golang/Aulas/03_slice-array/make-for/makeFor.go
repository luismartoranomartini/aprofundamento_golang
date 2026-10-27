package main

import "fmt"

func main()  {

	lista := make([]int, 1)
	lista[0] = 5
	lista = append(lista, 2)
	lista = append(lista, 3)
	fmt.Printf("%T\n", lista)

	
	// for i := 0; i < len(lista); i++{
	// 	fmt.Println(lista[i])
	// }


	// tirar a média da listagem

	somaTotal := 0
	// media := 0
	
	for i := 0; i < len(lista); i++{
		somaTotal += lista[i]
	}
	// media = somaTotal / len(lista)
	fmt.Println("Soma total = ", somaTotal / len(lista))
	// fmt.Println("Média total = ", media)
	// i := 0
	// for c, v := range lista{
	// 	fmt.Println(c, v)
	// 	i++
	// }

}