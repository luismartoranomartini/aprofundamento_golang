package model

import (
	"time"
	"errors"
)

type Compras struct {
	Mercado string
	Data    time.Time
	Itens []ItensCompra
}

type ItensCompra struct {
	Nome string 
}

func NewCompra(mercado string, data time.Time, nomeDositens []string) (*Compras, error){

	if mercado == "" {
		return nil, errors.New("Mercado obrigatório")
	}


	var itens []ItensCompra

	for _, nome := range nomeDositens {
		itens = append(itens, ItensCompra{Nome: nome})

	}

	return &Compras{
		Mercado: mercado,
		Data:    data,
		Itens:   itens,
	}, nil 
} 