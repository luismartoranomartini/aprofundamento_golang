package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Usuario struct {
	Nome   string `json:"nome"`
	Email  string `json:"email,omitempty"`
	Idade  int    `json:"idade"`
	senha  string
	Oculto string `json:"-"`
}

func main() {
	t := reflect.TypeOf(Usuario{})
	f, _ := t.FieldByName("Nome")
	fmt.Println("tag bruta:", f.Tag)
	fmt.Println("json:", f.Tag.Get("json"))

	u := Usuario{Nome: "Ana", Idade: 30, senha: "x", Oculto: "y"}
	b, _ := json.Marshal(u)
	fmt.Println(string(b))
}
