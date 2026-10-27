package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "Luís",
		"sobrenome": "Martini",
	}

	fmt.Println(usuario["nome"])
	fmt.Println(usuario["sobrenome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Luís",
			"segundo":  "Martini",
		},
		"curso": {
			"curso":     "Engenharia",
			"faculdade": "USP",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "curso")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "gêmeos",
	}
	fmt.Println(usuario2)
}
