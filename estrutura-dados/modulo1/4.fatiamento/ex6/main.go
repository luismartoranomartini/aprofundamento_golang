package main

import "fmt"

func posicao(s, sub string) int {
	for i := 0; i+len(sub) < len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return i
		}
	}
	return -1
}

func main() {

	s := "banana"
	fmt.Println(posicao(s, "b"))
	fmt.Println(posicao(s, "o"))
	fmt.Println(posicao(s, "anan"))
	fmt.Println(posicao(s, "na"))
}
