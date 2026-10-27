package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo      string
	turboLigado bool
	velAtua     int
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {
	carro1 := ferrari{"F40", false, 0}
	carro1.ligarTurbo()

	// nível de interface
	var carro2 esportivo = &ferrari{"F40", false, 0}
	carro2.ligarTurbo()

	fmt.Println(carro1, "\n",carro2)
}