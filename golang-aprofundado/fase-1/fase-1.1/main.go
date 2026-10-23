package main

import "fmt"

type Ponto struct {
	X, Y int
	Nome string
}

func main() {
	var i int
	var f float64
	var b bool
	var s string
	var p *int
	var sl []int
	var m map[string]int
	var e error

	fmt.Println(i)
	fmt.Println(f)
	fmt.Println(b)
	fmt.Println(s)
	fmt.Println(p)
	fmt.Println(sl)
	fmt.Println(m)
	fmt.Println(e)

	fmt.Println("Zero Value composto")
	var pt Ponto
	var arr [3]int
	fmt.Printf("%+v\n %v\n", pt, arr)
}
