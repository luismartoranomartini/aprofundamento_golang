package main

import "fmt"

func contaMaior(m map[string]int, limite int) int {
	var maiores int

	for _, v := range m {
		if v > limite {
			fmt.Println("Valores encontrados:", v)
			maiores++
		}
	}
	return maiores
}

func main() {
	mapa := map[string]int{"a": 3, "b": 7, "c": 9}
	limite := 5

	fmt.Println(contaMaior(mapa, limite))
}
