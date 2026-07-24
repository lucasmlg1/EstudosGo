package main

import (
	"errors"
	"fmt"
)

func main() {
	// int definido pela quantidade de bits, int8, int16, int32, int64
	numero := 100000
	fmt.Println(numero)

	// também possui o uint, que é um int sem sinal, como o "-"
	var numero2 uint = 32000
	fmt.Println(numero2)

	//alias (rune = int32)
	var numero3 rune = 32000
	fmt.Println(numero3)

	//alias (byte = uint8)
	var numero4 byte = 129
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123.45
	fmt.Println(numeroReal2)

	//sempre aspas duplas para string
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	//o mais próximo de char que se tem em Go, retorna o número da letra da tabela ASCII
	char := 'B' // num 66 da tabela
	fmt.Println(char)

	// FIM STRINGS

	var texto int16
	fmt.Println(texto)

	// Booleanos
	var booleano1 bool = true
	fmt.Println(booleano1)

	var booleano2 bool = false
	fmt.Println(booleano2)

	// O erro em Go é um tipo, o tratamento de erro é diferente do comum.
	var erro error = errors.New("erro interno")
	fmt.Println(erro)
	// o valor 0 é <nil>, que é tipo o null
}
