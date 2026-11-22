package main

import "fmt"

func contaPalavras(palavras []string) map[string]int {
	contador := make(map[string]int)

	for _, v := range palavras {
		contador[v]++
	}

	return contador
}

func main() {
	lista := []string{"go", "java", "go", "c", "go"}
	resultado := contaPalavras(lista)

	fmt.Println("go:", resultado["go"])
	fmt.Println("java:", resultado["java"])
	fmt.Println("c:", resultado["c"])
	fmt.Println("rust:", resultado["rust"])
}
