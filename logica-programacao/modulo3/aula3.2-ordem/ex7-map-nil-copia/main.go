package main

import "fmt"

func main() {
	var n map[string]int
	voltas := 0
	for range n {
		voltas++
	}
	fmt.Println("voltas no nil:", voltas)

	m := map[string]int{"a": 1, "b": 2}
	for _, v := range m {
		v = v * 10
		_ = v
	}
	fmt.Println(m)

	for k, v := range m {
		m[k] = v * 10
	}
	fmt.Println(m)
}
