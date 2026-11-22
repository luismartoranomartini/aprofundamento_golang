package main

import "fmt"

func main() {
	estoque := map[string]int{"maçã": 5, "pera": 0}
	fmt.Println(estoque["maçã"])
	fmt.Println(estoque["uva"])
	fmt.Println(len(estoque))
}
