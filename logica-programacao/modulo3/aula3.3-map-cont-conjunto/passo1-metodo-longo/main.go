package main

import "fmt"

func main() {
	frutas := []string{"maçã", "pera", "maçã", "uva", "maçã", "pera"}

	caderno := make(map[string]int)

	for _, fruta := range frutas {
		_, existe := caderno[fruta]

		if existe {
			caderno[fruta] = caderno[fruta] + 1
		} else {
			caderno[fruta] = 1
		}
		fmt.Println("li", fruta, "->", caderno)
	}
}
