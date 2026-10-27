package main

import "fmt"

func main() {
	var opcao int

	fmt.Print("Digite uma opção:\n1) Cadastrar\n2) Listar\n3) Editar\n4) Sair\nDigite:")
	fmt.Scan(&opcao)

	switch opcao {
	case 1:
		fmt.Println("Cadastrar")
	case 2:
		fmt.Println("Listar")
	case 3:
		fmt.Println("Editar")
	case 4:
		fmt.Println("Sair")
	default:
		fmt.Println("Opção inválida")
	}
}
