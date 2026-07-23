package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("")
	var array1 [5]int
	array1[0] = -14 // POSIÇÃO 1
	fmt.Println(array1[3])
	array2 := [4]string{"teste", "array", "testando", "surskity"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5} //Tentativa de flexibilizar o array utilizando o ..., faz com que ele baseie o tamanho na quantidade
	// de elementos que são passados
	fmt.Println(len(array3))

	slice := []int{3, 4, 5, 6} // nao tem tamanho fixo, muda de acordo com a necessidade, tamanho dinamico, mas também é tipado, bem utilizado em Go.
	fmt.Println(slice)

	slice = append(slice, 10) // ADICIONAR NOVO ELEMENTO COM APPEND
	fmt.Println(slice)

	slice2 := array2[1:3] // pedaço de um array armazenado em um slice.
	fmt.Println(slice2)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	// Arrays Internos
	//make aloca um espaço na memória para determinada função no programa
	slice4 := make([]float32, 10, 15)
	// make ele cria um array de 15 posições e retorna um slice de 10 das 15 posições, faz uma referência de um array interno

	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
	// Caso a capacidade previa definida pelo slice fosse estourada, seria criado um novo array com a capacidade em dobro
	// Por exemplo estourasse a capacidade máxima de 15, seria criado um novo array de 30 elementos.

	slice5 := make([]float32, 5)
	fmt.Println(slice5)
	fmt.Println(len(slice5))
	fmt.Println(cap(slice5))
	slice5 = append(slice5, 10)
	fmt.Println(len(slice5))
	fmt.Println(cap(slice5))
	//Array = uma lista de tamanho fixo
	//Slice = uma lista de tamanho dinâmico

}
