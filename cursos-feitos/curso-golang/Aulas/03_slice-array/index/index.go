package main
import "fmt"

func main() {

	listaToda := []int{2, 10, 9, 4, 5, 8, 1, 3, 6, 7, 11}

	var segundaLista = listaToda[:3]
	var terceiraLista = listaToda[4:]
	var quartaLista = listaToda[0:3]
	var capacidade = listaToda[1:4:4]
	var ultimoItem = listaToda[len(listaToda) -1 :]

	fmt.Println("Primeira ", listaToda)
	fmt.Println("Segunda ", segundaLista)
	fmt.Println("Terceira ",terceiraLista)
	fmt.Println("Quarta ", quartaLista)
	fmt.Println("Cap ", capacidade)
	fmt.Println("Ultimo: ", ultimoItem)
}