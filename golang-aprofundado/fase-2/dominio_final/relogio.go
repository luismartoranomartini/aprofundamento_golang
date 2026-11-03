package main

import (
	"fmt"
	"time"
)

type Relogio interface {
	horaExata() time.Time
}

type RelogioReal struct{}

type Cumprimentador struct {
	relogio Relogio
}

func (r RelogioReal) horaExata() time.Time {
	return time.Now()
}

func (c *Cumprimentador) cumprimentar(nome string) string {
	hora := c.relogio.horaExata().Hour()

	var saudacao string

	if hora < 12 {
		saudacao = "Bom dia"
	} else {
		saudacao = "Boa tarde"
	}

	return fmt.Sprintf("%s, %s!", saudacao, nome)
}

type RelogioFalso struct {
	hora time.Time
}

func (r RelogioFalso) horaExata() time.Time {
	return r.hora
}

func main() {
	// Exemplo de teste com hora fixa às 15h
	horaTeste := time.Date(2026, 11, 3, 15, 0, 0, 0, time.UTC)
	relogioFalso := RelogioFalso{hora: horaTeste}
	cumprimentador := Cumprimentador{relogio: relogioFalso}

	fmt.Println(cumprimentador.cumprimentar("Ana")) // "Boa tarde, Ana!"
}
