package main

import "fmt"

func main() {
	frutas := []string{"maçã", "pera", "maçã", "uva", "maçã", "pera"}
	caderno := make(map[string]int)

	for _, fruta := range frutas {
		caderno[fruta]++
		fmt.Println("li", fruta, "->", caderno)
	}
}
