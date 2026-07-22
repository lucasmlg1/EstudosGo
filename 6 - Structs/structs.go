package main

import "fmt"

type endereco struct {
	logradouro string
	numero     uint8
}

// um struct pode estar dentro de outro struct tranquilamente
type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

func main() {
	fmt.Println("test")

	// Structs é uma coleção de campo, o mais próximo de uma classe na linguagem

	var u usuario
	u.nome = "Lucas"
	u.idade = 20
	fmt.Println(u.nome)
	fmt.Println(u.idade)

	enderecoExemplo := endereco{"Rua Jornalista AssisChateubriand", 24}

	// inferencia de tipo utilizando um struct
	usuario2 := usuario{"Davi", 21, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Pedro"}
	fmt.Println(usuario3)
}
