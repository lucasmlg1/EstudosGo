package main

// alias
import (
	"fmt"
	f "fmt"
)

func main() {
	fmt.Println("Estruturas de Controle")
	numero := 9

	if numero > 10 {
		f.Println("Maior que 10.")
	} else if numero < 10 {
		f.Println("Menor que 10")
	} else {
		f.Println("O número é 10.")
	}

	if outroNumero := numero; outroNumero > 0 {
		f.Println("Maior que 0: ")
	} else {
		f.Println("Menor que 0")
	}

}
