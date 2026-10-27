package main

import "fmt"

func main() {
	lista := []int{4, 9, 6, 7}
	lista = append(lista, 5)
	lista[5] = 8

	listaDeString := []string{"Golang", "C#", "Java"}
	listaDeString = append(listaDeString, "Ruby")

	fmt.Println("Tamanho da Lista: ", len(lista))
	fmt.Println("Lista: ", lista)
	fmt.Println("Lista pos1: ", lista[0])
	fmt.Println("Lista pos2: ", lista[1])
	fmt.Println("Lista pos3: ", lista[2])
	fmt.Println("Lista pos4: ", lista[3])
	fmt.Println("Lista pos5: ", lista[4])
	fmt.Println("Lista de String: ", listaDeString)
	fmt.Println("Tamanhho Lista de String: ", len(listaDeString))


}