package main

import "fmt"

func inverterMap(mapa map[string]int) map[int]string {
	mapaInvertido := make(map[int]string)

	for k, v := range mapa {
		mapaInvertido[v] = k
	}
	return mapaInvertido
}

func main() {
	entrada := map[string]int{"a": 1, "b": 2}
	fmt.Println(inverterMap(entrada))
}
