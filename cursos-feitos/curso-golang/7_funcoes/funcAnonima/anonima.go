package main

import "fmt"

func main() {

	f := func(txt string) string {
		// fmt.Println(txt
		return fmt.Sprintf("Recebido -> %s", txt)
	}("Passando parâmetros")

	fmt.Println(f)
}
