package main

import "fmt"

type Livro struct {
	Titulo  string
	Paginas int
}

func main() {
	livros := []Livro{{"Go", 300}, {"Algoritmos:", 450}, {"Dados", 120}}

	soma := 0
	for _, l := range livros {
		soma += l.Paginas
	}
	fmt.Println("soma:", soma)

	maior := livros[0]
	for _, l := range livros {
		if l.Paginas > maior.Paginas {
			maior = l
		}
	}
	fmt.Println("maior:", maior.Titulo, maior.Paginas)

	//armadilha
	for _, l := range livros {
		l.Paginas = 0
	}
	fmt.Println("depois da armadilha:", livros)

	//correto: usar o índice para alterar o item do slice

	for i := range livros {
		livros[i].Paginas += 10
	}
	fmt.Println("depois do índice:", livros)
}
