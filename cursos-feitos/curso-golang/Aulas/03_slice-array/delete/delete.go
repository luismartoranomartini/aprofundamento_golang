package main

import "fmt"

func main() {

	m := make(map[string]int)
	m["SP"] = 100000
	m["CG"] = 90000
	m["RJ"] = 100000
	delete(m, "RJ")

	for c, v := range m {
		fmt.Printf("Cidade: %s | Prefeito: %d\n", c, v)

	}


}
