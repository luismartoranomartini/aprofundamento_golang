package main

import "fmt"

func main() {
	fmt.Println("abc" == "abc")    // true
	fmt.Println("abc" < "abd")     // true  ('c'=99 < 'd'=100)
	fmt.Println("Z" < "a")         // true  (90 < 97)
	fmt.Println("casa" < "casaco") // true  (prefixo menor vem antes)
	fmt.Println("" < "a")          // true
}
