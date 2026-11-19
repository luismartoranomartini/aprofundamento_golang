package main

import "fmt"

func main() {
	linhas, colunas := 3, 4

	m := make([][]int, linhas)
	fmt.Println(m, m[0] == nil, len(m[0]))

	for i := range m {
		m[i] = make([]int, colunas)
	}

	for i := 0; i < linhas; i++ {
		for j := 0; j < colunas; j++ {
			m[i][j] = i * j
		}
	}
	for _, linha := range m {
		fmt.Println(linha)
	}
}
