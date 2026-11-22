package main

import "fmt"

func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}

	primeiras := map[string]bool{}
	for i := 0; i < 200; i++ {
		for k := range m {
			primeiras[k] = true // anota a chave que abriu este loop
			break               // sai do loop: só a primeira interessa
		}
	}
	fmt.Println("chaves diferentes que abriram algum loop:", len(primeiras) > 1)
}
