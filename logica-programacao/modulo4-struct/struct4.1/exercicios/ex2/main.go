package main

import (
	"fmt"
)

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

func main() {
	livro := criarLivro("Go", 300, 50)
	fmt.Println(livro)

	livroComDesconto := aplicaDesconto(livro, 10)
	fmt.Println(livroComDesconto)
}
