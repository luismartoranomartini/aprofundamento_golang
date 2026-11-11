package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {

	}
	fmt.Println(i) // i não foi declarado
}

/* Previsão: Erro de compilação (undefined: i).

Por quê: Quando você declara i := 0 no cabeçalho do for, o escopo (tempo de vida) dessa variável existe apenas dentro do bloco { } do laço. Assim que o for termina, i é destruída e o fmt.Println(i) fora dele não a encontra. */
