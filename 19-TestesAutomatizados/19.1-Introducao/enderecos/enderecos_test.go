// TESTE DE UNIDADE
package enderecos_test

import (
	. "inroducao-testes/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

// Test + Nome da Função
func TestTipoDeEndereco(t *testing.T) {
	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida 123", "Avenida"},
		{"Estrada 987", "Estrada"},
		{"Rodovia do Parque", "Rodovia"},
		{"Praça Linda", "Tipo inválido"},
		{"RUA DOIS", "Rua"},
		{"avenida fogo", "Avenida"},
		{"", "Tipo inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)

		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				retornoRecebido,
				cenario.retornoEsperado,
			)
		}
	}
}
