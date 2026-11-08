package main

import "fmt"

type Animal struct {
	Nome string
}

func (a Animal) Falar() string {
	return a.Nome + " faz um som"
}

type Cachorro struct {
	Animal
	Raca string
}

func main() {
	c := Cachorro{Animal: Animal{Nome: "Rex"}, Raca: "vira-lata"}
	fmt.Println(c.Nome)
	fmt.Println(c.Animal.Nome)
	fmt.Println(c.Falar())
	c.Nome = "Totó"
	fmt.Println(c.Animal.Nome)
	fmt.Printf("%+v\n", c)
}
