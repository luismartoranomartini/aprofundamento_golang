package main

import "fmt"

func lerComExiste(mapa map[string]int, chave string) (int, bool) {
	if v, existe := mapa[chave]; existe {
		return v, existe
	}
	return 0, false
}

func main() {
	mapa := map[string]int{"a": 0, "b": 5}
	fmt.Println(lerComExiste(mapa, "z"))
}
