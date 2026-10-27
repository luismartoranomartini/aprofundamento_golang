package main

import "fmt"

func main() {

	m := make(map[string]string)
	m["SP"] = "Ricardo Nunes"
	m["CG"] = "Abilio"
	m["RJ"] = "Não sei"

	for c, v := range m {
		fmt.Printf("Cidade: %s | Prefeito: %s\n", c, v)

	}

}
