package main

import "fmt"

func main() {
	estoque := map[string]int{"maçã": 5, "pera": 0, "uva": 0, "kiwi": 2}

	for k, v := range estoque {
		if v == 0 {
			delete(estoque, k)
		}
	}
	fmt.Println(estoque, len(estoque))
}
