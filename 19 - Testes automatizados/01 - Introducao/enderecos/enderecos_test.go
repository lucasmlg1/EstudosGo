// teste unitário, de unidade (testar uma pequena parte do seu código)

package enderecos

import "testing"

func TestTipoDeEndereco(t *testing.T) {
	// Começa com TestXxxxXxxxX o nome da função obrigatoriamente.
	enderecoParaTeste := "Rua ABC"
	tipoDeEnderecoEsperado := "Rua"
	tipoDeEnderecoRecebido := TipoEndereco(enderecoParaTeste)
	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
		t.Error("O tipo é diferente do esperado!")
	}

}
