package main

import "fmt"

func maisCara(precos map[string]int) string {
	nome := ""
	maior := 0
	for k, v := range precos {
		if v > maior {
			maior = v
			nome = k
		}
	}
	return nome
}

func maisBarata(precos map[string]int) string {
	nome := ""
	menor := 10
	for k, v := range precos {
		if v < menor {
			menor = v
			nome = k
		}
	}
	return nome
}

func main() {
	precos := map[string]int{"pão": 8, "leite": 6, "queijo": 30, "ovo": 12}
	fmt.Println(maisCara(precos))
	fmt.Println(maisBarata(precos))
	fmt.Println(maisCara(nil) == "")
}
