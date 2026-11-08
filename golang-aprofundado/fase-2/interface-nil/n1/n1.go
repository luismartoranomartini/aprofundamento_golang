package main

import "fmt"

type MeuErro struct {
	Msg string
}

func (e *MeuErro) Error() string {
	return "meu erro: " + e.Msg
}

func main() {
	var p *MeuErro // ponteiro nil
	fmt.Println("p == nil:", p == nil)

	var e error // interface vazia
	fmt.Println("e == nil:", e == nil)
	fmt.Printf("e: %T %v\n", e, e)

	e = p // interface recebe um ponteiro nil
	fmt.Println("p == nil:", p == nil)
	fmt.Println("e == nil:", e == nil)
	fmt.Printf("e: %T %v\n", e, e)
}
