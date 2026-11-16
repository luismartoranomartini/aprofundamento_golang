package main

import "fmt"

func temRepetida(s string) bool {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[i] == runes[j] {
				return true
			}
		}
	}

	return false
}

func main() {
	fmt.Println(temRepetida("banana"))
	fmt.Println(temRepetida("gato"))
	fmt.Println(temRepetida(""))
	fmt.Println(temRepetida("ãé"))
}
