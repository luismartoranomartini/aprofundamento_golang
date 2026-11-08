package main

import "fmt"

func todasNaoRepetidas(s string) []rune {
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}
	var resultado []rune
	for _, r := range s {
		if m[r] == 1 {
			resultado = append(resultado, r)
		}
	}
	return resultado
}

func main() {
	palavras := "Adele"
	unicas := todasNaoRepetidas(palavras)
	if len(unicas) == 0 {
		fmt.Println("nenhuma se repete")
		return
	}
	fmt.Printf("Letras que não se repetem: ")
	for _, r := range unicas {
		fmt.Printf("\t%c\n", r)

	}
	fmt.Println()

}
