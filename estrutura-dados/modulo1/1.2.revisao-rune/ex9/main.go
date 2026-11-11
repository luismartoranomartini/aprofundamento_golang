package main

import "fmt"

func main() {
	r := []rune("ação")
	for i := 0; i < len(r)/2; i++ {
		j := len(r) - 1 - i
		r[i], r[j] = r[j], r[i]
	}
	fmt.Println(string(r))

	b := []byte("ação")
	for i := 0; i < len(b)/2; i++ {
		j := len(b) - 1 - i
		b[i], b[j] = b[j], b[i]
	}
	fmt.Printf("%q", string(b))
}
