package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1.")
}
func funcao2() {
	fmt.Println("Executando a função 2 ")
}

func alunoAprovado(n1, n2 float32) bool {
	if ((n1 + n2) / 2) >= 7 {
		return true
	}
	return false

}

func main() {
	// DEFER == ADIAR A EXECUÇÃO
	resultado := alunoAprovado(0, 5)
	defer funcao1()
	fmt.Println(resultado)
	funcao2()

}
