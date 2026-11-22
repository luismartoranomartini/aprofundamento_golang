package main

import "fmt"

func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	delete(m, "b")
	fmt.Println(m, len(m))

	delete(m, "zzz")
	fmt.Println(m, len(m))

	var n map[string]int
	delete(n, "a")
	fmt.Println(n, len(n))
}
