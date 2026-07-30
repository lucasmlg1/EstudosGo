package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoEndereco("Avenida Santo Amaro")
	fmt.Println(tipoEndereco)
}
