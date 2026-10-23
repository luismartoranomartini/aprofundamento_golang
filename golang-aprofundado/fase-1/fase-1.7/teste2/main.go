package main

import "fmt"

func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	visitados := 0
	for k := range m {
		visitados++
		if k == "a" {
			delete(m, "b") // remove "b" quando visita "a"
		}
	}
	fmt.Println(visitados)
}
