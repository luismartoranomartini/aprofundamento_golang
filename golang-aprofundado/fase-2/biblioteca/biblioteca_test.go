package main

import (
	"errors"
	"testing"
)

func TestEmprestarDisponivel(t *testing.T) {
	repo := &repoFalso{map[int]Livro{1: Livro{ID: 1, Disponivel: true}}}
	lib := &Biblioteca{livros: repo}

	err := lib.emprestar(1)
	if err != nil {
		t.Fatalf("erro de empréstimo: %v", err)
	}
	if repo.livros[1].Disponivel == true {
		t.Fatalf("erro de disponibilidade: %v", repo)
	}
}

func TestEmprestarIndisponivel(t *testing.T) {
	var alvo *ErrLivroIndisponivel

	repo := &repoFalso{map[int]Livro{1: Livro{ID: 1, Disponivel: false}}}
	lib := &Biblioteca{livros: repo}

	err := lib.emprestar(1)

	if !errors.As(err, &alvo) {
		t.Fatalf("erro: %v", err)
	}
}

func TestEmprestarInexistente(t *testing.T) {
	repo := &repoFalso{map[int]Livro{1: Livro{ID: 1, Disponivel: false}}}
	lib := &Biblioteca{livros: repo}

	err := lib.emprestar(1)
	if err == nil {
		t.Fatalf("erro: %v", err)
	}

	// err = lib.emprestar(10)
	// if err != nil {
	// 	t.Fatalf("erro: %v", err)
	// }
}
