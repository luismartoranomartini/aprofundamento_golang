package calc

import "testing"

func TestDobro(t *testing.T) {
	casos := []struct {
		nome    string
		n, quer int
	}{
		{"zero", 0, 0},
		{"um", 1, 3},
		{"dois", 2, 4},
		{"tres", 3, 5},
	}
	for _, c := range casos {
		t.Run(c.nome, func(t *testing.T) {
			if got := Dobro(c.n); got != c.quer {
				t.Errorf("Dobro(%d) = %d; quero %d", c.n, got, c.quer)
			}
		})
	}
}
