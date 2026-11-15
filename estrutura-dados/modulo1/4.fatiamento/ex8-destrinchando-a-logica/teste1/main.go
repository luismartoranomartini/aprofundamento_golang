package main

import "fmt"

func main() {
	s := "Luís Martini"
	fim := -1

	for i := len(s) - 1; i >= 0; i-- {
		fim = i + 1
		if s[i] == ' ' && fim == -1 {
			fmt.Println(s[i])
			fim = i + 1
		}
		if s[i] == ' ' && fim != -1 {
			fmt.Println(s[i+1 : fim])
		}
	}
}
