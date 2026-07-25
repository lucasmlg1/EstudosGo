package main

import "fmt"

// tudo atende a essa interface, é meio gambiarra na real
//é uma funcao generica que aceita um valor generico
func generica(inter interface{}) {
	fmt.Println(inter)
}

func main() {

	generica("String")
	generica(1)
	generica('2')

	//Gambiarra total kkkk
	mapa := map[interface{}]interface{}{
		1:             "Numero 1",
		float32(1000): true,
		true:          "String",
	}

	fmt.Println(mapa)

}
