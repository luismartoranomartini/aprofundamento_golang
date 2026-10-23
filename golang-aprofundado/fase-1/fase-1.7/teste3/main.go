package main

import "fmt"

func main() {
	m := map[int]bool{1: true, 2: true, 3: true}
	visitados := 0
	for k := range m {
		visitados++
		if k < 100 {
			m[k+100] = true // cria 101, 102, 103
		}
	}
	fmt.Println(visitados)
}
