package main

import "fmt"

func main() {
	// chave//valor
	fmt.Println("Maps")
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Marcelino",
		"cidade":    "Patos",
	}
	fmt.Println(usuario)
	// para acessar atributos específicos seria pela chave ele retorna o valor caso exista
	fmt.Println(usuario["nome"])
	fmt.Println(usuario["sobrenome"])
	fmt.Println(usuario["cidade"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro":  "Lucas",
			"sobrenome": "Marcelino",
		},
		"curso": {
			"nome": "Ciência da Computação",
		},
	}

	fmt.Println(usuario2["curso"])

	usuario2["teste"] = map[string]string{
		"nome": "teste",
	}

	fmt.Println(usuario2["teste"])

}
