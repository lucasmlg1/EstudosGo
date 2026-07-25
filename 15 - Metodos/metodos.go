package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) maiorDeIdade() {
	if u.idade >= 18 {
		fmt.Printf("O usuario %s é maior de idade com %d anos", u.nome, u.idade)
	} else {
		fmt.Printf("O usuario %s é menor de idade com %d anos", u.nome, u.idade)
	}

}

func (u usuario) salvar() {
	fmt.Println("Dentro do metodo salvar.")
	fmt.Println(u.nome)
	fmt.Println(u.idade)
}

func main() {
	usuario := usuario{"Lucas", 19}
	usuario.salvar()
	usuario.maiorDeIdade()
}
