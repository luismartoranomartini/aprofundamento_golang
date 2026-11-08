package main

import "fmt"

func main() {
	d := "7"[0]
	fmt.Println(d - '0')

	c := byte('g')
	fmt.Printf("%c\n", c-'a'+'A')
}
