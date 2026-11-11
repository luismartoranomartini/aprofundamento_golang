package main

import "fmt"

func main() {
	t := "ação"
	fmt.Printf("%q\n", t[0:2])
	fmt.Printf("%q\n", t[0:3])

	r := []rune(t)
	fmt.Println(string(r[:2]))
}
