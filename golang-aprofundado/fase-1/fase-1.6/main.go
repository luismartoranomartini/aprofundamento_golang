package main

import "fmt"

func main() {
	fmt.Println("Três formas de criar um map")
	idades := map[string]int{"ana": 30, "bia": 25}
	notas := make(map[string]float64)
	cache := make(map[int]string, 100)
	_ = cache
	fmt.Println(idades)
	fmt.Println(notas)
	fmt.Println(cache)

	notas["go"] = 9.5
	fmt.Println(idades["ana"], notas["go"])

	fmt.Println("Chave ausente NÃO dá erro: devolve o zero value do tipo do valor.")
	fmt.Println(idades["zeca"])

	fmt.Println("comma-ok distingue ausente de presente com valor zero.")
	v, ok := idades["zeca"]
	fmt.Println(v, ok)
	idades["zeca"] = 0
	v, ok = idades["zeca"]
	fmt.Println(v, ok)

	fmt.Println(idades)

	for v, k := range idades {
		fmt.Println(v, k)
	}

	fmt.Println("delete e len.")
	delete(idades, "zeca")
	fmt.Println(len(idades))

	fmt.Println("Map nil: lê, mas não escreve.")
	var nulo map[string]int
	fmt.Println(nulo["x"], len(nulo), nulo == nil)
	delete(nulo, "x")
	fmt.Println(nulo)

	a := map[string]int{"x": 1}
	b := a
	b["x"] = 99
	fmt.Println(a["x"])

	m := map[string]int{"a": 0}
	_, ok1 := m["a"]
	_, ok2 := m["b"]
	fmt.Println(ok1, ok2)
}
