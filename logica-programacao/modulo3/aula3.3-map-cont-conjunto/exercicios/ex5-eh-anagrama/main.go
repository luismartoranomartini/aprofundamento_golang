package main

import "fmt"

func ehAnagrama(a, b string) bool {
	contador := make(map[rune]int)

	for _, r := range a {
		contador[r]++
	}

	for _, r := range b {
		contador[r]--
	}

	for _, v := range contador {
		if v != 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(ehAnagrama("amor", "roma")) // true
	fmt.Println(ehAnagrama("casa", "caso")) // false: a fica com 1, o fica com -1
	fmt.Println(ehAnagrama("", ""))         // true: nada para comparar
	fmt.Println(ehAnagrama("aab", "abb"))   // false: a fica com 1, b fica com -1
}
