package main

import "fmt"

func main() {
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2
	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25
	soma1 := numero1 + numero2

	fmt.Println(soma1)
	// FIM ARITMETICOS

	//ATRIBUIÇÃO
	var variavel string = "String exemplo"
	variavel2 := "String exemplo2"
	fmt.Println(variavel, variavel2)

	//OPERADORES
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	fmt.Println("-------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	fmt.Println("---------------")
	//Operadores unarios

	numero := 10
	numero++
	numero += 15 // numero = numero + 15 (+=)
	fmt.Println(numero)

	numero10 := numero

	numero10--
	numero10 -= 15
	fmt.Println(numero10)

	//Fim dos operadores unarios

	// Adendo: não existe operadores ternarios em Go, condições como estas seriam feitas com o bom e velho if, else if, else

}
