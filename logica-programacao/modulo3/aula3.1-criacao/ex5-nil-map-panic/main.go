package main

import "fmt"

func main() {
	var m map[string]int

	fmt.Println(m["a"], len(m))
	v, ok := m["a"]
	fmt.Println(v, ok)

	m["a"] = 1
	fmt.Println(m)
}
