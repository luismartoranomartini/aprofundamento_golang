package main

import "fmt"

func main() {
	fmt.Println("a\tb")
	fmt.Println("\u00e9")
	fmt.Println("\x41")
	fmt.Println(`a\tb`)
	fmt.Println(`linha1
linha2`)
	fmt.Println(`C:\pasta\arquivo`)
	fmt.Printf("%q %q\n", "a\t", `a\tb`)
	//runes
	fmt.Println('\n')
	fmt.Println(' ')
	fmt.Println("Desafio")
	a := "x\ty"
	b := `x\ty`
	fmt.Println(a, b)
	fmt.Println(len(a), len(b))
}
