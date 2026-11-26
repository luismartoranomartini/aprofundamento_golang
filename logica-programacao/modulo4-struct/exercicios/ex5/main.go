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

func maisLivrosAutor(livros []Livro) (Autor, int) {
	contagem := make(map[Autor]int)
	var melhorAutor Autor
	maiorQtd := 0

	for _, l := range livros {
		contagem[l.Autor]++
	}

	for autor, qtd := range contagem {
		if qtd > maiorQtd {
			maiorQtd = qtd
			melhorAutor = autor
		}

	}

	return melhorAutor, maiorQtd

}

func main() {
	livrosEx5 := []Livro{
		{Titulo: "Go", Autor: Autor{Nome: "Ana", Pais: "Brasil"}},
		{Titulo: "Algoritmos", Autor: Autor{Nome: "Leo", Pais: "Portugal"}},
		{Titulo: "Dados", Autor: Autor{Nome: "Mia", Pais: "Brasil"}},
		{Titulo: "Redes", Autor: Autor{Nome: "Zeca", Pais: "Portugal"}},
		{Titulo: "Web", Autor: Autor{Nome: "Bia", Pais: "Chile"}},
	}
	mapa := paisLivros(livrosEx5)
	fmt.Println(mapa)

	livrosEx6 := []Livro{
		{Titulo: "Go", Autor: Autor{Nome: "Ana", Pais: "Brasil"}},
		{Titulo: "Dados", Autor: Autor{Nome: "Ana", Pais: "Portugal"}},
		{Titulo: "Redes", Autor: Autor{Nome: "Leo", Pais: "Brasil"}},
		{Titulo: "Web", Autor: Autor{Nome: "Ana", Pais: "Brasil"}},
	}
	maisLivros, qtd := maisLivrosAutor(livrosEx6)
	fmt.Println(maisLivros, qtd)
}
