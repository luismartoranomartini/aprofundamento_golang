package main

import "fmt"

func contem(s, sub string) bool {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}

func main() {
	s := "banana"

	fmt.Println(contem(s, "ana"))
	fmt.Println(contem(s, "nas"))
	fmt.Println(contem(s, "ban"))
	fmt.Println(contem(s, "nan"))
}
