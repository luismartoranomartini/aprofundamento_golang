package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct para Json

	p1 := produto{1, "Notebook", 1899.99, []string{"Promoção,", "Eletrônico"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))
	
	// json para struct
	var p2 produto
	jsonString := `{"id": 2, "nome": "Caneta", "preco": 8.9, "tags":["Papelaria,", "Importado"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2)

	var p3 produto
	jString := `{"id": 2, "nome": "Seringa", "preco": 1.0, "tags":["Farmácia,", "Comum"]}`
	json.Unmarshal([]byte(jString), &p3)
	fmt.Println("P3: ", p3)

	
}