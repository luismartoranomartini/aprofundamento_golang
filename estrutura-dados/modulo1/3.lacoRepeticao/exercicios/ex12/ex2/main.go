package main

import "fmt"

func main() {
	for i := 0; i < 3; {
		fmt.Println(i) // repete indefinidamente
	}
}

/* revisão: Loop infinito imprimindo 0 sem parar até você apertar Ctrl+C (e terminar com signal: interrupt).

Por quê: O pós-comando (onde normalmente vai o i++) ficou vazio. Como i começa em 0 e nunca é incrementado, a condição 0 < 3 permanece verdadeira para sempre. */
