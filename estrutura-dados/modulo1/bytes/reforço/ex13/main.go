package main

import "fmt"

func trocaPontas(s string) string {
	if s == "" {
		return s
	}
	b := []byte(s)

	b[0], b[len(b)-1] = b[len(b)-1], b[0]
	return string(b)
}

func main() {
	fmt.Println(trocaPontas("abcdef"))
	fmt.Println(trocaPontas("abc"))
	fmt.Println(trocaPontas("x"))
	fmt.Printf("%q\n", trocaPontas(""))
}
