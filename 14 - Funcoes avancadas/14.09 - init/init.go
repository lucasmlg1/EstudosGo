package main

import "fmt"

// Função que é executada antes mesmo da main
func init() {
	fmt.Println("Executando funcao init")
}

func main() {

	fmt.Println("Função main sendo executada")
}
