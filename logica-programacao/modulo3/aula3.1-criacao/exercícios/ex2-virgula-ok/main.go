package main

import "fmt"

func main() {
	estoque := map[string]int{"maçã": 5, "pera": 0}
	v, ok := estoque["maçã"]
	fmt.Println(v, ok)

	v, ok = estoque["pera"]
	fmt.Println(v, ok)

	v, ok = estoque["uva"]
	fmt.Println(v, ok)
}
