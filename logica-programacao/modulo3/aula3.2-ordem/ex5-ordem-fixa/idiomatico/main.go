package main

import (
	"fmt"
	"maps"
	"slices"
)

func imprimirEmOrdem(notas map[string]int) {
	for _, nome := range slices.Sorted(maps.Keys(notas)) {
		fmt.Println(nome, notas[nome])
	}
}

func main() {
	notas := map[string]int{"Caio": 6, "Ana": 8, "Bia": 10}
	imprimirEmOrdem(notas)
}
