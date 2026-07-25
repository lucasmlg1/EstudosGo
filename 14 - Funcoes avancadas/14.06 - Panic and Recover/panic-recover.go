package main

import "fmt"

func recuperarExecução() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso")
	}
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperarExecução()
	media := (n1 + n2) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	// o programa entra em panico e interrompe o fluxo de execução caso nao tenha um recover junto
	// Antes de "matar" o programa ele vai chamar todas as clausulas que tem defer
	panic("A MEDIA É EXATAMENTE 6")

}

func main() {

	result := alunoEstaAprovado(6, 6)
	fmt.Println(result)
	fmt.Println("Teste de sequência de execução")

}
