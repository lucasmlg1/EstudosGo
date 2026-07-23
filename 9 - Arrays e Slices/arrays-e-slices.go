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
	slice = append(slice, 10)
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))
}
