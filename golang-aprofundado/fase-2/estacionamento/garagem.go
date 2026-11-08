package main

import (
	"errors"
	"fmt"
	"time"
)

type Vaga struct {
	ID      int
	Ocupada bool
}

type Veiculo struct {
	Placa string
}

type Ticket struct {
	Local     Vaga
	CarroMoto Veiculo
	Hora      time.Time
}

type Garagem struct {
	entrada RepositorioVagas
}

func (g *Garagem) entrar(placa string) (Ticket, error) {
	vaga, err := g.entrada.buscar()
	if err != nil {
		return Ticket{}, ErrSemVaga{Placa: placa}
	}
	vaga.Ocupada = true

	err = g.entrada.salvar(vaga)
	if err != nil {
		return Ticket{}, err
	}
	return Ticket{Local: vaga, CarroMoto: Veiculo{Placa: placa}, Hora: time.Now()}, nil
}

type ErrSemVaga struct {
	Placa string
}

func (ev ErrSemVaga) Error() string {
	return fmt.Sprintf("Vaga não encontrada: %v", ev.Placa)
}

type RepositorioVagas interface {
	buscar() (Vaga, error)
	salvar(v Vaga) error
}

type RepoFalso struct {
	vaga map[int]Vaga
}

func (rf *RepoFalso) salvar(v Vaga) error {
	rf.vaga[v.ID] = v
	return nil
}

func (rf *RepoFalso) buscar() (Vaga, error) {
	for _, v := range rf.vaga {
		if !v.Ocupada {
			return v, nil
		}
	}
	return Vaga{}, errors.New("vaga não encontrada")
}

func main() {

}
