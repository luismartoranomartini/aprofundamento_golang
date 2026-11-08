package main

import "fmt"

func naoRepete(s string) int {
outer:
	for i, r := range s {
		for j, rr := range s {
			if i != j && r == rr {
				continue outer
			}
		}
		return i

	}
	return -1
}

func main() {
	l := "Maa"
	fmt.Println(naoRepete(l))
}
