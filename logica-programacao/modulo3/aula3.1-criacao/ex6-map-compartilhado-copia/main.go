package main

import "fmt"

func marca(m map[string]int) {
	m["Zé"] = 7
}

func main() {
	a := map[string]int{"Ana": 8}
	marca(a)
	fmt.Println(a)

	b := a
	b["Ana"] = 0
	fmt.Println(a, b)

	// Para copiar de verdade: map novo + um par por vez (o range sobre map é a Aula 3.2).
	c := make(map[string]int)
	for k, v := range a {
		c[k] = v
	}

	c["Ana"] = 99
	fmt.Println(a, c)

}
