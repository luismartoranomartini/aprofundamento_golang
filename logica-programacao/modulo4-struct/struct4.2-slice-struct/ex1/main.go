package main

import "fmt"

// Livro guarda o título e a quantidade de páginas
type Livro struct {
	Titulo  string
	Paginas int
}

func main() {
	// forma 1: var sem valor inicial: slice nil, sem nenhum item
	var a []Livro
	fmt.Println("a:", len(a), a == nil)

	// forma 2: make cria n itens JÁ existentes, todos com o valor zero de Livro
	b := make([]Livro, 2)
	fmt.Printf("b: %+v\n", b)

	b[0] = Livro{"Go", 300}            // preenche a posição 0
	b = append(b, Livro{"Dados", 120}) // append acrescenta DEPOIS dos 2 itens zerados
	fmt.Printf("b: %+v | len %d\n", b, len(b))

	// forma 3: literal; o nome Livro pode ser omitido dentro das chaves
	c := []Livro{{"Redes", 200}, {"Web", 150}}
	fmt.Println("c:", c, len(c))

	// append funciona também em slice nil
	a = append(a, Livro{"Algoritmos", 450})
	fmt.Println("a depois do append:", a, len(a), a == nil)
}
