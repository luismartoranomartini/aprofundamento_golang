package main

import "fmt"

type Base struct{ ID int }

func (b Base) Nome() string    { return "base" }
func (b Base) Mostrar() string { return "mostrar: " + b.Nome() }

type Item struct{ Base }

func (i Item) Nome() string { return "item" }

func main() {
	it := Item{Base{7}}
	fmt.Println(it.Nome(), it.Base.Nome())
	fmt.Println(it.Mostrar(), it.ID)
	it.ID = 8
	fmt.Println(it.Base.ID)
}
