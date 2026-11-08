package main

import "fmt"

func soma(a, b int) int { return a + b }

func aplica(f func(int, int) int, x, y int) int {
	return f(x, y)
}

func main() {
	f := soma // função como valor
	fmt.Printf("%T\n", f)
	fmt.Println(f(1, 2))
	fmt.Println(aplica(soma, 3, 4))

	// função anônima chamada na hora
	r := func(x int) int { return x * x }(5)
	fmt.Println(r)

	// func literal guardada em variável
	dobro := func(x int) int { return 2 * x }
	fmt.Println(dobro(8))

	var nada func()
	fmt.Println(nada == nil)
}
