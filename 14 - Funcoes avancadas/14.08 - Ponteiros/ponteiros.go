package main

import "fmt"

func inverterSinal(n1 int) int {
	return n1 * -1
}

func inverterSinalPonteiro(n1 *int) {
	*n1 = *n1 * -1
	fmt.Println(*n1)
}

func main() {

	numero := 20
	fmt.Println(numero)
	numeroNovo := inverterSinal(numero)
	fmt.Println(numeroNovo)
	fmt.Println(numero)

	novoValor := 30
	inverterSinalPonteiro(&novoValor)
	fmt.Println(novoValor)

}
