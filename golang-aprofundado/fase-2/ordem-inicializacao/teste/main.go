package main

import "fmt"

var x = pr("x", y+1)
var y = pr("y", 10)
var z = pr("z", 5)

func pr(nome string, v int) int {
	fmt.Println("var", nome)
	return v
}

func init() { fmt.Println("init 1", x, y, z) }
func init() { fmt.Println("init 2") }

func main() { fmt.Println("main") }
