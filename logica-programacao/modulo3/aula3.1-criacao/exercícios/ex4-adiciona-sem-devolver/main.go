package main

import "fmt"

func adiciona(m map[string]int, k string, v int) {
	if _, ok := m[k]; !ok {
		m[k] = v
	}
}

func main() {
	m := make(map[string]int)

	adiciona(m, "a", 0)
	fmt.Println(m)

	adiciona(m, "a", 99)
	fmt.Println(m)

	/*
		Não, a função não precisa devolver o map.

	Em Go, tudo é passado por valor (cópia), mas um map é, por baixo dos panos, um ponteiro para a estrutura interna da tabela hash na memória (hmap).
	*/

}
