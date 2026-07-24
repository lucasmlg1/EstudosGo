package main

import "fmt"

// sem clausula break :O

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Segunda-feira"
		fallthrough
	case 2:
		return "Terça-feira"
	case 3:
		return "Quarta-feira"
	case 4:
		return ("Quinta-feira")
	case 5:
		return ("Sexta-feira")
	case 6:
		return ("Sábado")
	case 7:
		return ("Domingo")
	default:
		return "Insira algo valido"
	}
}

func main() {
	fmt.Println("Switch Case")
	fmt.Println(diaDaSemana(3))
}
