package main

import (
	"fmt"
	"math"
)

type geometria interface {
	area() float64
}

type retangulo struct {
	largura, altura float64
}

func (r retangulo) area() float64 {
	return r.largura * r.altura
}

type circulo struct {
	radius float64
}

func (c circulo) area() float64 {
	return math.Pi * c.radius * c.radius
}

func ExibirGeometria(g geometria) {
	fmt.Println(g.area())

}

type ProblemaNetwork struct {
	rede     bool
	hardware bool
}

func ExibeErro(err error) {
	fmt.Println(err.Error())
}
func (p ProblemaNetwork) Error() string {
	if p.rede {
		return "Problema de rede"
	} else if p.hardware {
		return "Problema de hardware"
	} else {
		return "outro problema"
	}
}
func main() {
	fmt.Println("Iniciando...")


	var lista []interface{} 
	lista = append(lista, 10)
	lista = append(lista, 7.5)
	lista = append(lista, "teste")
	fmt.Println(lista)

	for _, valor := range lista{

		if v, ok := valor.(string); ok {
			fmt.Println(v, " string")
		} else {
			fmt.Println(v)
		}

	}

}
