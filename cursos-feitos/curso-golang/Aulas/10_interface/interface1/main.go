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

func main() {
	fmt.Println("Iniciando...")

	retangulo := retangulo{
		largura: 2,
		altura:  3,
	}

	circulo := circulo{
		radius: 3,
	}
	
	ExibirGeometria(retangulo)
	ExibirGeometria(circulo)

}
