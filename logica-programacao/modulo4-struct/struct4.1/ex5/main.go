package main

import "fmt"

type Ponto struct {
	X, Y int
}

func main() {
	a := Ponto{1, 2}
	b := Ponto{1, 2}
	c := Ponto{2, 1}

	fmt.Println("a == b", a == b)
	fmt.Println("a == c", a == c)

	lugares := make(map[Ponto]string)
	lugares[Ponto{1, 2}] = "casa"
	lugares[Ponto{3, 4}] = "escola"

	fmt.Println("lugares[a]:", lugares[a])
	_, ok := lugares[Ponto{9, 9}]
	fmt.Println("{9, 9} existe?", ok)
	fmt.Println("len:", len(lugares))
}
