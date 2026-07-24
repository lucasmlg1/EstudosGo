package main

import (
	"fmt"
	"time"
)

func main() {
	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println(i)
	// 	time.Sleep(time.Second)
	// }

	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementando J: ", j)
	// 	time.Sleep(time.Second)
	// }

	nomes := [3]string{"João", "Davi", "Lucas"}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, letra)
	}

	usuario := map[string]string{
		"nome":      "Lucas",
		"sobrenome": "Marcelino",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// NAO EH POSSIVEL USAR DO FOR RANGE EM UMA ESTRUTURA TIPO O STRUCT, APENAS EM MAP E SLICE/ARRAY
	// type usuarioStruct struct {
	// 	nome      string
	// 	sobrenome string
	// }

	// for chave, valor = range usuario3 {
	// 	fmt.Println(chave, valor)
	// }

	//Como se fosse um while true, continua executando até que uma condição de parada seja exigida.
	for true {
		fmt.Println("EXECUTANDO INFINITAMENTE")
		time.Sleep(time.Second)
	}

}
