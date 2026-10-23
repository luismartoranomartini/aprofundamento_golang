package main

import "fmt"

type Ponto struct{ X, Y int }
type Reta struct{ A, B Ponto }

func main() {
	r1 := Reta{Ponto{0, 0}, Ponto{1, 1}}
	r2 := r1
	r2.B.X = 5
	p := &r1.A
	p.Y = 9
	fmt.Println(r1 == r2, r1, r2)
	r2.B.X = 1
	r2.A.Y = 9
	fmt.Println(r1 == r2)
}
