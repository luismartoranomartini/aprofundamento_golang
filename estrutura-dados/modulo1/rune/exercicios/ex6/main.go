package main

import "fmt"

func main() {
	r := []byte("maçã")
	for i := 0; i < len(r)/2; i++ {
		j := len(r) - 1 - i
		r[i], r[j] = r[j], r[i]
	}
	fmt.Printf("%q\n", string(r))
}
