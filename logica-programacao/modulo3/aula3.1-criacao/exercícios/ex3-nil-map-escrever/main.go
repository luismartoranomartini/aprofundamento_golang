package main

import "fmt"

func main() {
	var m map[string]int
	fmt.Println(m["x"])
	// Escrever em um nil map causa um runtime panic:
	// "panic: assignment to entry in nil map"
	// Por quê? Porque o map ainda não tem nenhuma tabela hash alocada na memória
	// para armazenar pares chave-valor.
	m = make(map[string]int)

	m["x"] = 1

	fmt.Println(m["x"])

}
