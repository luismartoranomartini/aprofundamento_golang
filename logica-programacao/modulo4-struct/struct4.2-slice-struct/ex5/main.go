package main

import "fmt"

type Livro struct {
	Titulo  string
	Paginas int
}

func removerNoLugar(livros []Livro, i int) []Livro {
	return append(livros[:i], livros[i+1:]...)
}

func removerCopia(livros []Livro, i int) []Livro {
	novo := make([]Livro, 0, len(livros)-1)
	novo = append(novo, livros[:i]...)
	novo = append(novo, livros[i+1:]...)
	return novo
}

func nova() []Livro {
	return []Livro{{"A", 1}, {"B", 2}, {"C", 3}, {"D", 4}}
}

func main() {
	for _, i := range []int{0, 2, 3} {
		o := nova()
		r := removerCopia(o, i)
		fmt.Println("copia i =", i, "->", r, "| original:", o)
	}

	for _, i := range []int{0, 2, 3} {
		o := nova()
		r := removerNoLugar(o, i)
		fmt.Println("lugar i=", i, "->", r, "| original:", o)
	}
}
