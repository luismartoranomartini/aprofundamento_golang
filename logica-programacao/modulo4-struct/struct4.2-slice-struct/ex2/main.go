package main

import "fmt"

// Livro guarda o título e a quantidade de páginas
type Livro struct {
	Titulo  string
	Paginas int
}

// buscar procura o primeiro livro com o título dado; o bool diz se achou
func buscar(livros []Livro, titulo string) (Livro, bool) {
	for _, l := range livros { // l é cópia: só lemos
		if l.Titulo == titulo {
			return l, true // achou: devolve o livro e true
		}
	}
	return Livro{}, false // não achou (ou slice vazio): valor zero e false
}

func main() {
	livros := []Livro{{"Go", 300}, {"Dados", 120}, {"Go", 999}}

	l, ok := buscar(livros, "Dados") // existe
	fmt.Println(l, ok)

	l, ok = buscar(livros, "Redes") // não existe
	fmt.Println(l, ok)

	l, ok = buscar(livros, "Go") // título repetido: devolve o primeiro
	fmt.Println(l, ok)

	l, ok = buscar(nil, "Go") // slice vazio
	fmt.Println(l, ok)
}
