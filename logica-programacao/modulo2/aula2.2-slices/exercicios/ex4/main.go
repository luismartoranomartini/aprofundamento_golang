package main

import "fmt"

func zera(s [3]int) {
	for i := range s {
		s[i] = 0
	}
	fmt.Println(s)
}

func main() {
	numeros := [3]int{1, 2, 3}
	zera(numeros)
	fmt.Println(numeros)
}

/* Nada muda no original ([1 2 3] permanece intacto).

Diferente de um slice — que é apenas um header contendo um ponteiro para dados externos —, um array em Go é um tipo de valor puro (como um int ou uma struct). A variável numeros contém diretamente os três inteiros alocados em sequência na memória. */
