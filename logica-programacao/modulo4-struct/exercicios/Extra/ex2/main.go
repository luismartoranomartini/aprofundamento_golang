package main

import "fmt"

type Nota struct {
	Aluno      string
	Disciplina string
	Valor      float64
}

func notaAlunos(notas []Nota) (string, float64) { // nome do aluno e média
	somaAlunos := make(map[string]float64)
	qtdProvas := make(map[string]int)

	for _, n := range notas {
		somaAlunos[n.Aluno] += n.Valor
		qtdProvas[n.Aluno]++
	}

	var vencedor string
	var maiorMedia float64

	for aluno, soma := range somaAlunos {
		media := soma / float64(qtdProvas[aluno])
		if media > maiorMedia {
			vencedor = aluno
			maiorMedia = media
		}
	}

	return vencedor, maiorMedia
}

func main() {
	notas := []Nota{
		{Aluno: "Ana", Disciplina: "Go", Valor: 8.0},
		{Aluno: "Leo", Disciplina: "Go", Valor: 7.0},
		{Aluno: "Bia", Disciplina: "Go", Valor: 10.0},
		{Aluno: "Leo", Disciplina: "Redes", Valor: 7.0},
		{Aluno: "Ana", Disciplina: "Redes", Valor: 9.0},
		{Aluno: "Leo", Disciplina: "Dados", Valor: 7.0},
		{Aluno: "Bia", Disciplina: "Redes", Valor: 8.0},
	}

	aluno, media := notaAlunos(notas)
	fmt.Println(aluno, media)
}
