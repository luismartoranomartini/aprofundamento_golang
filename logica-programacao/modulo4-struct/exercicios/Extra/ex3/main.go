package main

import (
	"fmt"
)

type Nota struct {
	Aluno      string
	Disciplina string
	Valor      float64
}

type Resumo struct {
	Soma      float64
	QtdProvas int
}

func notaAlunos(notas []Nota) (string, float64) { // nome do aluno e média
	resultadoTotal := make(map[string]Resumo)

	for _, nota := range notas {
		temp := resultadoTotal[nota.Aluno]
		temp.Soma += nota.Valor
		temp.QtdProvas++
		resultadoTotal[nota.Aluno] = temp
	}
	var nome string
	var maiorMedia float64
	var recorde float64

	for aluno, resumo := range resultadoTotal {
		maiorMedia = resumo.Soma / float64(resumo.QtdProvas)
		if maiorMedia > recorde {
			recorde = maiorMedia
			nome = aluno
		}
	}
	return nome, recorde
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
