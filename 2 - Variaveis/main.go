package main

import "fmt"

func main() {
	// Fortemente tipado, todas as variáveis tem um tipo
	var variavel1 string = "Variável 1" // declaração explícita
	variavel2 := "Variável 2"           // declaração implícita -> inferência da tipo
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "variavel 3"
		variavel4 string = "variavel 4"
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "variavel 5", "variavel 6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante"
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
