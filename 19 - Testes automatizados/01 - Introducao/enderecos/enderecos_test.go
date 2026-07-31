// teste unitário, de unidade (testar uma pequena parte do seu código)

package enderecos_test

import (
	. "introducao-testes/enderecos"
	"testing"
)

type cenarioTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	// Começa com TestXxxxXxxxX o nome da função obrigatoriamente
	t.Parallel()
	cenariosDeTeste := []cenarioTeste{
		{"Avenida ABC", "Avenida"},
		{"Rua das rosas", "Rua"},
		{"Rodovia dos imigrantes", "Rodovia"},
		{"Praça das rosas", "Tipo inválido de endereco."},
		{"Estrada qualquer", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		{"", "Tipo inválido de endereco."},
	}

	for _, cenario := range cenariosDeTeste {
		tipoEnderecoRecebido := TipoEndereco(cenario.enderecoInserido)
		if tipoEnderecoRecebido != cenario.retornoEsperado {
			t.Error("O tipo é diferente do esperado!")
		}
	}

}

func TestQualquerTipo(t *testing.T) {
	t.Parallel()
	if 1 > 2 {
		t.Error("Teste quebrou")
	}
}
