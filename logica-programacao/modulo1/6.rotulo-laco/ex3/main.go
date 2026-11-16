package main

import "fmt"

func contem(s, sub string) bool {
janela:
	for i := 0; i+len(sub) < len(s); i++ {
		for j := 0; j < len(sub); j++ {
			if s[i+j] != sub[j] {
				continue janela
			}
		}
		return true
	}
	return false
}

func main() {
	fmt.Println(contem("banana", "nan"))
	fmt.Println(contem("banana", "nas"))
	fmt.Println(contem("banana", "ba"))
}
