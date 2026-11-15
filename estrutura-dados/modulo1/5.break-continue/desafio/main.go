package main

import "fmt"

func contaSemDigito(s string) int {
	contaStr := 0
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		if runes[i] >= '0' && runes[i] <= '9' {
			continue
		}
		contaStr++
	}

	return contaStr
}

func main() {
	fmt.Println(contaSemDigito("abc7d9"))
	fmt.Println(contaSemDigito("123"))
	fmt.Println(contaSemDigito(""))

	fmt.Println("=== Caso de borda ===")
	fmt.Println(contaSemDigito("Luís"))
	fmt.Println(contaSemDigito("ação7"))
	fmt.Println(contaSemDigito("7ação"))
	fmt.Println(contaSemDigito("e1e"))

}
