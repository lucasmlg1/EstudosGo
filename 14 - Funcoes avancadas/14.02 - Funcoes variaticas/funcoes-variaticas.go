package main

import "fmt"

// vai receber de 1 a n ints
func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total

}

//Função com um parametro fixo e outro com elementos variaticos(sempre precisa estar no final da função e só pode ter 1 por função.)
func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	fmt.Println(soma(1, 2, 3, 4, 5, 6, 7))

	escrever("Teste", 1, 2, 3, 4)

}
