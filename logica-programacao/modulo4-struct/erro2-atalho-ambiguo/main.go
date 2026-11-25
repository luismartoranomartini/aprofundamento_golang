package main

import "fmt"

type A struct{ Nome string }
type B struct{ Nome string }

type C struct {
	A
	B
}

func main() {
	c := C{A{"X"}, B{"Y"}}
	fmt.Println(c.A.Nome, c.B.Nome)
	fmt.Println(c.Nome)
}
