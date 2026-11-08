package calc

import (
	"errors"
	"testing"
)

func TestDividirSimples(t *testing.T) {
	got, err := Dividir(10, 2)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}
	if got != 5 {
		t.Errorf("Dividir(10, 2) = %d; quero 5", got)
	}
}

func TestDividir(t *testing.T) {
	casos := []struct {
		nome    string
		a, b    int
		quer    int
		querErr error
	}{
		{"exata", 10, 2, 5, nil},
		{"truncada", 7, 2, 3, nil},
		{"negativa", -9, 3, -3, nil},
		{"por zero", 1, 0, 0, ErrDivisaoPorZero},
	}
	for _, c := range casos {
		t.Run(c.nome, func(t *testing.T) {
			got, err := Dividir(c.a, c.b)
			if !errors.Is(err, c.querErr) {
				t.Fatalf("erro = %v; quero %v", err, c.querErr)
			}
			if got != c.quer {
				t.Errorf("Dividir(%d, %d) = %d; quero %d", c.a, c.b, got, c.quer)
			}
		})
	}
}
