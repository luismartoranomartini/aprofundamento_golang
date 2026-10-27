package main

import (
	"interface/pkg/conversor"
	"interface/pkg/conversor/json"
	"interface/pkg/conversor/xml"
)

func main() {
	gerenciador := conversor.NovoGerenciador()
	gerenciador.RegistrarConversor(json.ConversorJSON{})
	gerenciador.RegistrarConversor(xml.ConversorXML{})

	dadosJSON := `{
		"pessoas": {
			"nome": "Luís",
			"idade": 49
		}
	}`

	dadosXML := `
	<pessoa>
		<nome>Luís</nome>
		<idade>49</idade>
	</pessoa>
	`

	gerenciador.Processar("JSON", dadosJSON)
	gerenciador.Processar("XML", dadosXML)
}
