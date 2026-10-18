package main

import (
	"fmt"
	"maps"
	"slices"
	"sort"
)

func main() {
	m := map[string]int{"banana": 2, "maçã": 5, "uva": 1, "pera": 7}

	chaves := make([]string, 0, len(m))

	for k := range m {
		chaves = append(chaves, k)
	}
	sort.Strings(chaves)
	for _, k := range chaves {
		fmt.Println(k, m[k])
	}

	for _, k := range slices.Sorted(maps.Keys(m)) {
		fmt.Print(k, " ")
	}
	fmt.Println()

	contagem := map[string]int{"a": 1, "b": 2}
	for _, v := range contagem {
		v = v * 100
	}
	fmt.Println(contagem)

	for k, v := range contagem {
		contagem[k] = v * 100
	}
	fmt.Println(contagem)

	preco := map[string]int{"a": 10, "b": 20}
	for _, p := range preco {
		p = p + 1
	}
	for k := range preco {
		preco[k] += 5
	}
	fmt.Println(preco)

}
