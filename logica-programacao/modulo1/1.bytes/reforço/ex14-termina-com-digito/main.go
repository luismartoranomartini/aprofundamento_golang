package main

import "fmt"

func terminaComDigito(s string) bool {
	if s == "" {
		return false
	}

	b := s[len(s)-1]
	return b >= '0' && b <= '9'
}

func main() {
	fmt.Println(terminaComDigito("abc7"))
	fmt.Println(terminaComDigito("7abc"))
	fmt.Println(terminaComDigito("9"))
	fmt.Println(terminaComDigito("abc"))
	fmt.Printf("%t\n", terminaComDigito(""))
}
