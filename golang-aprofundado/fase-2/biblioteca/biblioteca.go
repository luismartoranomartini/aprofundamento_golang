package main

import (
	"errors"
	"fmt"
	"time"
)

type ErrLivroIndisponivel struct {
	ID int
}

func (e *ErrLivroIndisponivel) Error() string {
	return fmt.Sprintf("livro ID %d falhou", e.ID)
}

type Livro struct {
	ID         int
	Titulo     string
	Disponivel bool
}

type Leitor struct {
	ID   int
	Nome string
}

type Emprestimo struct {
	ID     int
	Livro  Livro
	Leitor Leitor
	Data   time.Time
}

type Acervo struct {
	Livros []Livro
}

type Biblioteca struct {
	livros RepositorioLivros
}

func (b *Biblioteca) emprestar(id int) error {
	livro, err := b.livros.buscar(id)
	if err != nil {
		return err
	}
	if !livro.Disponivel {
		return &ErrLivroIndisponivel{ID: id}
	}

	livro.Disponivel = false

	err = b.livros.salvar(livro)
	if err != nil {
		return err
	}
	return nil
}

type repoFalso struct {
	livros map[int]Livro
}

func (rf *repoFalso) buscar(id int) (Livro, error) {
	livro, ok := rf.livros[id]
	if !ok {
		return Livro{}, errors.New("Livro não existe")
	}

	return livro, nil
}

func (rf *repoFalso) salvar(l Livro) error {
	rf.livros[l.ID] = l

	return nil
}

type RepositorioLivros interface {
	buscar(id int) (Livro, error)
	salvar(l Livro) error
}

func main() {

}
