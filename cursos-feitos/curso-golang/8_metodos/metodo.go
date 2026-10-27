package main

import "fmt"

type usuario struct {
	nome  string
	idade uint
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário '%s' no banco de dados\n", u.nome)
}

func (u usuario) maiorIdade() bool {
	return u.idade >= 18
}

func escrever() {
	fmt.Println("escrevend0")
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"User 1", 20}
	fmt.Println(usuario1)
	usuario1.salvar()
	maiorIdade := usuario1.maiorIdade()
	fmt.Println(maiorIdade)

	usuario2 := usuario{"Davi", 16}
	usuario2.salvar()
	usuario2.fazerAniversario()
	usuario2.fazerAniversario()
	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}
