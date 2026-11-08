package main

import "fmt"

type Animal struct{ Nome string }
type Dono struct{ *Animal }

func main() {
	d := Dono{}
	fmt.Println(d.Animal == nil)
	fmt.Println(d.Nome)
}
