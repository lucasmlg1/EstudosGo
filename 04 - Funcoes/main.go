package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// uma função consegue ter 2 retornos (Muito usado!)
func calculos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	fmt.Println(somar(3, 4))
	// variavel sendo atribuida de uma funcao
	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}
	resultado := f("Go é interessante")
	fmt.Println(resultado)

	resultadoSoma, resultadoSubtracao := calculos(10, 15)
	fmt.Println("Soma: ", resultadoSoma, "|| Subtracao: ", resultadoSubtracao)
}
