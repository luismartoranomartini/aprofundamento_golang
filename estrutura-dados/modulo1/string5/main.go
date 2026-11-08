package main

import "fmt"

func main() {
	s := "gato"
	// s[0] = 'p'

	b := []byte(s)
	b[0] = 'p'
	s2 := string(b)
	fmt.Println(s, s2)

}
