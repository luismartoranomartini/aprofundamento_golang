package main

import "fmt"

type Ponto struct {
	X, Y int
}

type Segmento struct {
	De, Ate Ponto
	Nome    string
}

func main() {
	a := Ponto{1, 2}
	b := Ponto{Y: 5, X: 4}
	c := Ponto{X: 9}
	d := &Ponto{X: 7, Y: 8}
	fmt.Println(a, b, c, d)

	s := Segmento{De: Ponto{0, 0}, Ate: Ponto{2, 4}, Nome: "s1"}
	s.Ate.X = 10
	fmt.Printf("%+v\n", s)

	cfg := struct {
		Host  string
		Porta int
	}{"localhost", 8080}
	fmt.Println(cfg.Host, cfg.Porta)
	fmt.Printf("%T\n", cfg)
}
