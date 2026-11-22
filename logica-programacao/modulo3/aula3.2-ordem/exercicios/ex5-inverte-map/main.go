package main

import "fmt"

func inverte(m map[string]int) map[int]string {
	mapaInvertido := make(map[int]string)
	for k, v := range m {
		mapaInvertido[v] = k
	}
	return mapaInvertido
}

func main() {
	m1 := map[string]int{"a": 1, "b": 2}
	fmt.Println("Teste 1:", inverte(m1))

	m2 := map[string]int{"a": 1, "b": 1}
	fmt.Println("Teste 2:", inverte(m2))

	/*
		Por que o resultado do Teste 2 pode mudar a cada execução?

		Em Go, a ordem de iteração de um map com "for range" é intencionalmente
		randomizada pelo runtime. Como "a" e "b" possuem o mesmo valor (1), ambos
		tentam gravar na mesma chave no mapa invertido (invertido[1]).

		A chave que for visitada por último no loop sobrescreve a anterior:
		- Se o loop passar por "a" e depois "b", o resultado será map[1:b].
		- Se o loop passar por "b" e depois "a", o resultado será map[1:a].
	*/
}
