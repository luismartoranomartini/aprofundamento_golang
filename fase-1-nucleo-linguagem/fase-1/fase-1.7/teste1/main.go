package main

import "fmt"

func main() {
	m := map[int]bool{1: true, 2: true, 3: true, 4: true}
	visitados := 0
	for k := range m {
		visitados++
		for outro := range m {
			if outro != k {
				delete(m, outro) // apaga todas as outras chaves
			}
		}
	}
	fmt.Println(visitados, len(m)) // 1 1
}
