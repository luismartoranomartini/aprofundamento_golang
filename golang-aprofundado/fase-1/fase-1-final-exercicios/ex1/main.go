package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func main() {
	texto := "o gato viu o rato e o rato viu o gato"

	limpo := strings.Map(func(r rune) rune {
		if unicode.IsPunct(r) {
			return ' '
		}
		return unicode.ToLower(r)
	}, texto)

	palavras := strings.Fields(limpo)

	contagem := make(map[string]int)
	for _, p := range palavras {
		contagem[p]++
	}

	chaves := make([]string, 0, len(contagem))
	for p := range contagem {
		chaves = append(chaves, p)
	}
	sort.Strings(chaves)

	for _, p := range chaves {
		fmt.Printf("%s %d\n", p, contagem[p])
	}
}
