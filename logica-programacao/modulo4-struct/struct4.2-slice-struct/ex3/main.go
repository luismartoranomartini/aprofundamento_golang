package main

import "fmt"

// Livro guarda o título e a quantidade de páginas
type Livro struct {
	Titulo  string
	Paginas int
}

// maioresQue devolve um slice NOVO só com os livros de mais de n páginas
func maioresQue(livros []Livro, n int) []Livro {
	var resultado []Livro // começa nil
	for _, l := range livros {
		if l.Paginas > n {
			resultado = append(resultado, l) // copia o livro para o resultado
		}
	}
	return resultado
}

func main() {
	livros := []Livro{{"Go", 300}, {"Dados", 120}, {"Redes", 200}}

	r := maioresQue(livros, 150) // alguns
	fmt.Println(r, len(r))

	r = maioresQue(livros, 1000) // nenhum
	fmt.Println(r, len(r), r == nil)

	r = maioresQue(livros, 0) // todos
	fmt.Println(r, len(r))

	r = maioresQue(nil, 0) // slice vazio de entrada
	fmt.Println(r, len(r))
}
