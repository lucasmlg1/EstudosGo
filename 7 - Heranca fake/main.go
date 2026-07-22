package main

import "fmt"

// todo estudante é uma pessoa também
type estudante struct {
	pessoa
	curso     string
	faculdade string
}

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

func main() {
	p1 := pessoa{"Pedro", "Batista", 22, 179}
	e1 := estudante{p1, "CC", "UFMG"}
	fmt.Println(e1)
	fmt.Println("Nome:", e1.nome)
	fmt.Println("Altura:", e1.altura)

}
