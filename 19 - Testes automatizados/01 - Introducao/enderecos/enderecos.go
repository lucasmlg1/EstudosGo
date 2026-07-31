package enderecos

import "strings"

//TipoEndereco verifica se tem um endereco valido e o retorna
func TipoEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	enderecoComLetrasMinuscula := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoComLetrasMinuscula, " ")[0]

	enderecoValido := false
	for _, tipo := range tiposValidos {
		if primeiraPalavraDoEndereco == tipo {
			enderecoValido = true
		}
	}
	if enderecoValido {
		return strings.Title(primeiraPalavraDoEndereco)
	}
	return "Tipo inválido de endereco."
}
