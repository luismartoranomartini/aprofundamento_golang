package main

import "fmt"

func main() {
	a := "C:\\temp\\novo"
	fmt.Println(a)
	b := `C:\temp\novo`
	fmt.Println(b)
	fmt.Println(a == b)
	fmt.Println(len(b))
	fmt.Println(len("a\tb"))
}
