package main

import (
	"fmt"
	"time"
)

func main() {
	// Concorrência != Paralelismo!
	// go routines são funções que podem ser chamadas para executar um outro método sem aquele fluxo padrão de execução de funções,
	// esperando uma acabar para começar outra.
	go escrever("Olá, mundo")

	escrever("Programando em Go!")
}

func escrever(texto string) {
	for true {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
