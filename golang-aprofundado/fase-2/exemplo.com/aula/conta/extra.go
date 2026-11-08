package conta

func Dobrar(c *Conta) {
	c.saldo = c.saldo * 2
	_ = validar(1)
}
