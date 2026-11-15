package main

import "fmt"

func ultimaPalavra(s string) string {
	fim := -1

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' && fim == -1 {
			fim = i + 1
		}
		if s[i] == ' ' && fim != -1 {
			return s[i+1 : fim]
		}
	}
	if fim != -1 {
		return s[:fim]
	}
	return ""
}

func main() {
	fmt.Println(ultimaPalavra("Luís Martini"))
	fmt.Println(ultimaPalavra(" ir casa"))
	fmt.Println(ultimaPalavra("um"))
	fmt.Println(ultimaPalavra(""))

}
