package main

import "fmt"

func copia(m map[string]int) map[string]int {

	novo := map[string]int{}

	for k, v := range m {
		novo[k] = v
	}

	return novo
}

func main() {
	mapa := make(map[string]int)
	mapa["Ana"] = 10

	c := copia(mapa)
	c["Bia"] = 60
	fmt.Println(mapa)
	fmt.Println(c)
}
