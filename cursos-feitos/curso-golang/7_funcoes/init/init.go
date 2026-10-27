package main

import "fmt"

var n int

func init() {
	fmt.Println("Excutando a init")
	n = 10
}

func main() {
	fmt.Println("Main sendo executada")
	fmt.Println(n)

}
