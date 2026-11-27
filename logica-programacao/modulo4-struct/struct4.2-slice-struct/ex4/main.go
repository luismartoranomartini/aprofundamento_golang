package main

import "fmt"

type Livro struct {
	Titulo  string
	Paginas int
}

func main() {
	livros := []Livro{{"Go", 300}, {"Dados", 120}}

	// 1) range com l: l é uma cópia; alterar l não altera o slice
	for _, l := range livros {
		l.Paginas = 0
	}
	fmt.Println("range com l:", livros)

	// 2) cópia de um item por índice: também é cópia
	x := livros[0]
	x.Paginas = 0
	fmt.Println("cópia x: ", livros)

	// 3) índice direto: altera o item do slice
	livros[0].Paginas = 301
	fmt.Println("índice:", livros)

	// 4) para alterar todos: laço pelo índice
	for i := range livros {
		livros[i].Paginas += 10
	}
	fmt.Println("todos +10:", livros)

	// 5) trocar o item inteiro também vale
	livros[1] = Livro{"Redes", 200}
	fmt.Println("troca:", livros)

}
