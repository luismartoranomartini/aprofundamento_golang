package main

import "fmt"

type MeuErro struct{ Msg string }

// receptor ponteiro que trata nil: funciona mesmo com ponteiro nil
func (e *MeuErro) Error() string {
	if e == nil {
		return "(erro nil)"
	}
	return "meu erro: " + e.Msg
}

type Outro struct{ Msg string }

// receptor valor: com ponteiro nil dentro da interface, dá panic
func (o Outro) Error() string { return "outro: " + o.Msg }

func main() {
	var p *MeuErro
	var e error = p
	fmt.Println(e.Error()) // funciona: o método trata nil

	var q *Outro
	var e2 error = q
	defer func() { fmt.Println("recuperado:", recover()) }()
	fmt.Println(e2.Error()) // panic
}
