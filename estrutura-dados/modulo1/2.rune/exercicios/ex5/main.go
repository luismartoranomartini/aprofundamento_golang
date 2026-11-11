package main

import "fmt"

func main() {
	str := []rune("coração")

	for i, r := range str {
		if r == 'ç' {
			str[i] = 'c'
		} else if r == 'ã' {
			str[i] = 'a'
		}
	}
	fmt.Println(string(str))
}
