package main

import (
	"fmt"
	"modulo/auxiliar"
)

func main() {
	fmt.Println("Escrevendo do arquivo main ")
	auxiliar.Escrever()

	// Go não é uma linguagem orientada a objeto, logo para sabermos se uma função pode ou não
	// ser exportada por outra classe(Não existe classes em Go!!), é preciso analisar a primeira letra da função.
	// minúscula -> apenas pelo mesmo pacote, não pode ser chamada em outras classes de pacotes diferentes.
	// maiuscula -> pode ser chamada por classes de pacotes diferentes.
}
