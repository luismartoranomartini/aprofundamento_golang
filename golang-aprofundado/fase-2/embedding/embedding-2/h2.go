package main

import "fmt"

type Animal struct {
	Nome string
}

func (a Animal) Falar() string      { return a.Nome + " faz um som" }
func (a Animal) Apresentar() string { return "Apresentando: " + a.Falar() }

// func recebe(a Animal)               {}

type Cachorro struct {
	Animal
}

func (c Cachorro) Falar() string { return c.Nome + " late" }

func main() {
	c := Cachorro{Animal{"Rex"}}
	fmt.Println(c.Falar())
	fmt.Println(c.Animal.Falar())
	fmt.Println(c.Apresentar())
	_ = Cachorro{Nome: "Rex"}
	// c1 := Cachorro{Animal{"Rex"}}
	// recebe(c)
}
