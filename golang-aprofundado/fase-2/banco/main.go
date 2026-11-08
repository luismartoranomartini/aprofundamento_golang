package main

import (
	"errors"
	"fmt"
)

type Cliente struct {
	ID   int
	Nome string
}

type Transacao struct {
	Tipo  string
	Valor int64
}

type Conta struct {
	ID        int
	Titular   Cliente
	saldo     int64
	historico []Transacao
}

type ErroValorInvalido struct {
	Valor int64
}

func (ev ErroValorInvalido) Error() string {
	return fmt.Sprintf("valor inválido: %d", ev.Valor)
}

type ErroSaldoInsuficiente struct {
	Saldo  int64
	Pedido int64
}

func (es ErroSaldoInsuficiente) Error() string {
	return fmt.Sprintf("saldo insuficiente: saldo %d, pedido %d", es.Saldo, es.Pedido)
}

func NovaConta(id int, titular Cliente) *Conta {
	conta := &Conta{
		ID:      id,
		Titular: titular,
	}
	return conta
}

func (c *Conta) Saldo() int64 {
	return c.saldo
}

func (c *Conta) Historico() []Transacao {
	return c.historico
}

func (c *Conta) Depositar(valor int64) error {
	if valor <= 0 {
		return ErroValorInvalido{Valor: valor}
	}
	c.saldo += valor
	c.historico = append(c.historico, Transacao{Tipo: "deposito", Valor: valor})
	return nil
}

func (c *Conta) Sacar(valor int64) error {
	if valor <= 0 {
		return ErroValorInvalido{Valor: valor}
	}
	if valor > c.saldo {
		return ErroSaldoInsuficiente{Saldo: c.saldo, Pedido: valor}
	}
	c.saldo -= valor
	c.historico = append(c.historico, Transacao{Tipo: "saque", Valor: valor})
	return nil
}

func main() {

	cli := Cliente{ID: 1, Nome: "Ana"}
	c := NovaConta(10, cli)
	fmt.Println(c.Depositar(10000))
	fmt.Println(c.Depositar(-5))
	fmt.Println(c.Sacar(3000))
	err := c.Sacar(20000)
	fmt.Println(err)
	var es ErroSaldoInsuficiente
	if errors.As(err, &es) {
		fmt.Println("faltam", es.Pedido-es.Saldo)
	}
	var ev ErroValorInvalido
	fmt.Println(errors.As(c.Sacar(0), &ev), ev.Valor)
	fmt.Println(c.Saldo(), len(c.Historico()))
	for _, t := range c.Historico() {
		fmt.Println(t)
	}
}
