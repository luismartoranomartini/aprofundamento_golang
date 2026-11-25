package main

import "fmt"

func main() {
	conjunto := make(map[string]struct{})

	conjunto["go"] = struct{}{}
	conjunto["map"] = struct{}{}
	conjunto["go"] = struct{}{}

	fmt.Println("tamanho", len(conjunto))

	_, existe := conjunto["map"]
	fmt.Println("map existe?", existe)

	_, existe = conjunto["rune"]
	fmt.Println("rune existe?", existe)

	delete(conjunto, "map")
	_, existe = conjunto["map"]
	fmt.Println("map existe depois do delete?", existe, "tamanho:", len(conjunto))
}
