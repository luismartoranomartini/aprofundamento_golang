package main

import (
	"testing"
	"time"
)

func TestCumprimentar(t *testing.T) {
	// Tabela de casos de teste
	testes := []struct {
		nomeCenario string
		hora        int
		nomePessoa  string
		esperado    string
	}{
		{
			nomeCenario: "Manhã cedo (3h)",
			hora:        3,
			nomePessoa:  "Ana",
			esperado:    "Bom dia, Ana!",
		},
		{
			nomeCenario: "Fim da manhã (11h)",
			hora:        11,
			nomePessoa:  "Carlos",
			esperado:    "Bom dia, Carlos!",
		},
		{
			nomeCenario: "Meio-dia exato (12h)",
			hora:        12,
			nomePessoa:  "Beatriz",
			esperado:    "Boa tarde, Beatriz!",
		},
		{
			nomeCenario: "Noite (20h)",
			hora:        20,
			nomePessoa:  "Ana",
			esperado:    "Boa tarde, Ana!",
		},
		{
			nomeCenario: "Quase meia-noite (23h)",
			hora:        23,
			nomePessoa:  "Lucas",
			esperado:    "Boa tarde, Lucas!",
		},
	}

	for _, tc := range testes {
		t.Run(tc.nomeCenario, func(t *testing.T) {
			// Cria o momento no tempo com a hora definida pelo caso de teste
			horaSimulada := time.Date(2026, 1, 1, tc.hora, 0, 0, 0, time.UTC)
			relogioFalso := RelogioFalso{hora: horaSimulada}
			cumprimentador := Cumprimentador{relogio: relogioFalso}

			obtido := cumprimentador.cumprimentar(tc.nomePessoa)

			if obtido != tc.esperado {
				t.Errorf("para às %02dh: obtido %q, esperado %q", tc.hora, obtido, tc.esperado)
			}
		})
	}
}
