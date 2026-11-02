package conta

import "fmt"

type Conta struct {
	Titular string
	saldo   float64
}

func Nova(titular string) *Conta {
	return &Conta{Titular: titular}
}

func (c *Conta) Depositar(v float64) {
	if validar(v) {
		c.saldo += v
	}
}

func (c *Conta) Saldo() float64 { return c.saldo }

func validar(v float64) bool { return v > 0 }

func Resumo(c *Conta) string {
	return fmt.Sprintf("%s: %.2f", c.Titular, c.saldo)
}
