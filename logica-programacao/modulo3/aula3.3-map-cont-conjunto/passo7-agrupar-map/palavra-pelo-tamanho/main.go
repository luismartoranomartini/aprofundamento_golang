package main

import "fmt"

func main() {
	grupos := make(map[int][]string)

	palavras := []string{"go", "map", "rune", "lei", "if", "byte", "for"}

	for _, p := range palavras {
		tam := len(p)
		grupos[tam] = append(grupos[tam], p)
	}

	fmt.Println("tam 9:", grupos[9], len(grupos[9]))

	for tam := 2; tam <= 4; tam++ {
		fmt.Println(tam, grupos[tam], "quantidade:", len(grupos[tam]))
	}
}
