package main

import "fmt"

func main() {
	texto := "olá"
	b := make([]byte, 3)
	n := copy(b, texto)
	c := "\xc3\xa3"

	fmt.Println("len(texto) em bytes:", len(texto))
	fmt.Println("copiados:", n)
	fmt.Printf("b = %v %%q = %q\n", b, b)
	fmt.Printf("UTF-8 = %s\n", c)
}
