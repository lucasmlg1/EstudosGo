package main

import "fmt"

func main() {
	// é uma função sem nome que pode ser atribuída a variáveis, executada imediatamente ou passada como argumento.

	//função anônima que é executada logo após ser declarada.
	func() {
		fmt.Println("Olá, mundo")
	}()

	func(texto string) {
		fmt.Println(texto)
	}("Passando parametro")

	//Em Go (Golang), a função fmt.Sprintf cria e retorna uma string formatada usando verbos de formatação (como %s ou %d)

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("E Passando parametro")

	fmt.Println(retorno)
}
