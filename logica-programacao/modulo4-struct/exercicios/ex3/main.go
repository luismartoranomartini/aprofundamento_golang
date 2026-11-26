package main

import "fmt"

type Livro struct {
	Titulo  string
	Paginas int
	Preco   float64
}

func criarLivro(titulo string, paginas int, preco float64) Livro {
	return Livro{
		Titulo:  titulo,
		Paginas: paginas,
		Preco:   preco,
	}
}

func aplicaDesconto(livro Livro, perc float64) Livro {
	livro.Preco -= livro.Preco * (perc / 100)
	return livro
}

func maisCaro(livros []Livro) Livro {
	if len(livros) == 0 {
		return Livro{}
	}
	caro := livros[0]

	for _, l := range livros {
		if l.Preco > caro.Preco {
			caro = l
		}
	}
	return caro
}

func somaPaginas(livros []Livro) int {
	soma := 0
	for _, l := range livros {
		soma += l.Paginas
	}
	return soma
}

func descontoEmTodos(livros []Livro, perc float64) {
	for i := range livros {
		livros[i] = aplicaDesconto(livros[i], perc)
		fmt.Println(livros[i].Preco)
	}
}

func main() {
	// livro := criarLivro("Go", 300, 50)
	// fmt.Println(livro)
	//
	// livroComDesconto := aplicaDesconto(livro, 10)
	// fmt.Println(livroComDesconto)

	livros := []Livro{
		{Titulo: "Go", Paginas: 300, Preco: 50},
		{Titulo: "Algoritmos", Paginas: 450, Preco: 80},
		{Titulo: "Dados", Paginas: 120, Preco: 35},
	}
	// caro := maisCaro(livros)
	// soma := somaPaginas(livros)
	// fmt.Println(caro.Titulo, caro.Preco, soma)

	descontoEmTodos(livros, 10)
	fmt.Println(livros[0].Preco, livros[1].Preco, livros[2].Preco)

	for _, l := range livros {
		fmt.Println(l.Preco)
	}

	// fmt.Println(maisCaro(nil), somaPaginas(nil))
}
