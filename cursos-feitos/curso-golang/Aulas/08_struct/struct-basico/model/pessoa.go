package model

import (
	"time"
)

type Pessoa struct {
	Nome           string
	Endereco       Endereco
	DataNascimento time.Time
	Idade          int
}

// Método
func (p *Pessoa) ProcessaIdade() {
	anoNascimento := p.DataNascimento.Year()
	anoAtual := time.Now().Year()
	p.Idade = anoAtual - anoNascimento
	// fmt.Println("Idade: ", p.Idade)
}

func (p Pessoa) IdadeAtual() int {

	anoNascimento := p.DataNascimento.Year()
	anoAtual := time.Now().Year()
	return anoAtual - anoNascimento
}

// Função
func CalculaIdade(p Pessoa) int {
	anoNascimento := p.DataNascimento.Year()
	anoAtual := time.Now().Year()
	return anoAtual - anoNascimento
}
