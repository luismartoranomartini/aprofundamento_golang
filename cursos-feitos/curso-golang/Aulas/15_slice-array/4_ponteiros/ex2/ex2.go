package main

import "fmt"

func main() {

	p := NovaPessoa("Luís", 25)
	fmt.Println(p)
	fmt.Println(p.Nome, p.Idade)
	fmt.Println(p.Telefone())
	p.AtualizarIdade(26)
	fmt.Println(p.Nome, p.Idade)
	p.AtualizarTelefone("38615920")
	fmt.Println(p.Telefone())
}

type Pessoa struct {
	Nome     string
	Idade    *int // 0
	telefone *string
}

func NovaPessoa(nome string, idade int) Pessoa {
	return Pessoa{
		Nome:  nome,
		Idade: &idade,
	}
}

func (p Pessoa) Telefone() string {
	if p.telefone == nil {
		return ""
	}
	return *p.telefone
}

func (p *Pessoa) AtualizarIdade(idade int) {
	p.Idade = &idade
}

func (p *Pessoa) AtualizarTelefone(telefone string) {
	p.telefone = &telefone
}
