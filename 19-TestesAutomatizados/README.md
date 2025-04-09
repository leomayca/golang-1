## Testes Automatizados em Go

O Go oferece uma ferramenta robusta e simples para criação de **testes automatizados** nativos na linguagem, permitindo escrever, executar e analisar testes com comandos diretos do terminal.

### Estrutura e Convenção dos Testes

#### 1. Nome do arquivo de teste

- Todo arquivo de teste deve terminar com `_test.go`.
- Exemplo:
  ```
  calculadora.go
  calculadora_test.go
  ```

#### 2. Nome da função de teste

- Deve começar com `Test` e ser seguido do nome da função a ser testada.
- A assinatura deve receber um `*testing.T`.

```go
package calculadora

func Somar(a, b int) int {
	return a + b
}
```

```go
package calculadora_test

import (
	"calculadora"
	"testing"
)

func TestSomar(t *testing.T) {
	resultado := calculadora.Somar(2, 3)

	if resultado != 5 {
		t.Errorf("Esperado 5, mas recebeu %d", resultado)
	}
}
```

### Comandos Básicos para Testes

#### 1. Executar testes

```bash
go test
```

Executa todos os testes do pacote atual.

#### 2. Executar testes com mais detalhes

```bash
go test -v
```

Mostra informações detalhadas sobre quais testes passaram ou falharam.

#### 3. Executar testes com cobertura

```bash
go test -cover
```

Mostra a porcentagem de código coberto por testes.

### Análise de Cobertura

#### 1. Gerar relatório de cobertura em arquivo

```bash
go test -coverprofile=coverage.out
```

Gera um arquivo `coverage.out` com os dados de cobertura.

#### 2. Visualizar o relatório de cobertura no terminal

```bash
go tool cover -func=coverage.out
```

#### 3. Abrir relatório visual no navegador

```bash
go tool cover -html=coverage.out
```

Gera um relatório em HTML e abre no navegador para facilitar a análise.

### Boas Práticas

- **Testes devem ser pequenos e independentes**.
- Sempre use **cenários com valores esperados e casos de borda**.
- Organize os testes por funcionalidades e módulos.
- Use `t.Run()` para subtestes organizados:

```go
func TestOperacoes(t *testing.T) {
	t.Run("Soma", func(t *testing.T) {
		resultado := Somar(1, 2)
		if resultado != 3 {
			t.Fail()
		}
	})

	t.Run("Subtração", func(t *testing.T) {
		resultado := Subtrair(5, 2)
		if resultado != 3 {
			t.Fail()
		}
	})
}
```

### Dica: Testes de Tabela (Table-Driven Tests)

Muito comum em Go:

```go
func TestDobro(t *testing.T) {
	cenarios := []struct {
		entrada int
		esperado int
	}{
		{2, 4},
		{5, 10},
		{0, 0},
	}

	for _, c := range cenarios {
		resultado := Dobro(c.entrada)
		if resultado != c.esperado {
			t.Errorf("Para entrada %d, esperado %d, recebido %d",
				c.entrada, c.esperado, resultado)
		}
	}
}
```
