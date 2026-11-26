package main

import "fmt"

type Livro struct {
	Titulo  string
	Paginas int
	Preco   float64
	Autor   Autor
}

type Autor struct {
	Nome string
	Pais string
}

func paisLivros(livros []Livro) map[string][]string {
	resultado := make(map[string][]string)

	for _, l := range livros {
		pais := l.Autor.Pais
		resultado[pais] = append(resultado[pais], l.Titulo)
	}
	return resultado
}

func main() {
	livros := []Livro{
		{Titulo: "Go", Autor: Autor{Nome: "Ana", Pais: "Brasil"}},
		{Titulo: "Algoritmos", Autor: Autor{Nome: "Leo", Pais: "Portugal"}},
		{Titulo: "Dados", Autor: Autor{Nome: "Mia", Pais: "Brasil"}},
		{Titulo: "Redes", Autor: Autor{Nome: "Zeca", Pais: "Portugal"}},
		{Titulo: "Web", Autor: Autor{Nome: "Bia", Pais: "Chile"}},
	}
	mapa := paisLivros(livros)
	fmt.Println(mapa)
}
