package main

import (
	"testing"
)

func TestEntrarComVaga(t *testing.T) {
	repo := &RepoFalso{map[int]Vaga{1: Vaga{ID: 1, Ocupada: false}}}
	g := &Garagem{entrada: repo}

	ticket, err := g.entrar("ABC1234")
	if err != nil {
		t.Fatalf("erro ao encontrar vaga: %v", err)
	}

	if ticket.CarroMoto.Placa != "ABC1234" {
		t.Fatal(ticket)
	}
	return

}
