package main

import "fmt"

func primeiraPalavra(s string) string {
	inicio := -1

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && inicio == -1 {
			inicio = i
		}
		if s[i] == ' ' && inicio != -1 {
			return s[inicio:i]

		}
	}
	if inicio != -1 {
		return s[inicio:]
	}

	return ""
}

func main() {
	fmt.Println(primeiraPalavra(" ir para casa"))
	fmt.Println(primeiraPalavra("um"))
	fmt.Println(primeiraPalavra(""))
	fmt.Println(primeiraPalavra("ab"))

}
