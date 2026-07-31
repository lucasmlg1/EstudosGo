package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"Idade"`
}

func main() {
	// Map para Json ou Struct para Json
	c := cachorro{"Malu", "York Shire", 3}
	fmt.Println(c)
	cachorroEmJson, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(bytes.NewBuffer(cachorroEmJson))

	c2 := map[string]string{
		"Nome":  "Kirby",
		"Raca":  "Pinscher",
		"Idade": "11",
	}

	cachorro2EmJSON, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))
}
