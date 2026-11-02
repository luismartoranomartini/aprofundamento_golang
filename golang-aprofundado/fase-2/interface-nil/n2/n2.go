package main

import "fmt"

type MeuErro struct{ Msg string }

func (e *MeuErro) Error() string {
	return "meu erro: " + e.Msg
}

// BUG: devolve a variável *MeuErro como error, mesmo quando ela é nil
func buscaRuim(achou bool) error {
	var err *MeuErro
	if !achou {
		err = &MeuErro{"não achou"}
	}
	return err
}

// CERTO: devolve nil explícito quando não há erro
func buscaBoa(achou bool) error {
	if !achou {
		return &MeuErro{"não achou"}
	}
	return nil
}

func main() {
	if err := buscaRuim(true); err != nil {
		fmt.Printf("ruim: achou, mas err != nil → %T %v\n", err, err == nil)
	}
	if err := buscaBoa(true); err != nil {
		fmt.Println("boa: não deveria imprimir")
	} else {
		fmt.Println("boa: err == nil")
	}
}
