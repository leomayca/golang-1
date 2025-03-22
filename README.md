# Golang

## 4. Pacotes

### O que é um Módulo?

Um módulo é um conjunto de pacotes que compõem um projeto em Go. Ele facilita a organização do código e a gestão de dependências.

### Criando um Módulo

Para criar um módulo, utilize o seguinte comando:

```sh
 go mod init [NOME_DO_MODULO]
```

Esse comando gerará um arquivo `go.mod`, que centraliza todas as dependências do projeto. Ele funciona de maneira semelhante ao `package.json` no Node.js.

### Compilação e Execução

Com o módulo configurado, é possível rodar o código de diferentes formas:

- **Execução direta**:

  ```sh
  go run main.go
  ```

  Executa o código diretamente sem gerar um executável.

- **Compilação e execução**:

  ```sh
  go build
  ```

  Esse comando cria um arquivo binário (executável) contendo todo o código compilado. O executável gerado pode ser executado diretamente no sistema operacional.

- **Instalação do binário**:

  ```sh
  go install
  ```

  O `go install` compila o código e move o binário gerado para o diretório `$GOPATH/bin`, tornando-o acessível globalmente no sistema, permitindo sua execução de qualquer lugar no terminal.

### Visibilidade de Funções e Variáveis

Em Go, a visibilidade de funções e variáveis dentro de um pacote segue uma regra simples:

- **Públicas**: Nomes que começam com letra maiúscula são exportados e podem ser acessados por outros pacotes.
- **Privadas**: Nomes que começam com letra minúscula são acessíveis apenas dentro do próprio pacote.

#### Exemplo:

```go
package exemplo

// Função pública
func Publica() {
    fmt.Println("Esta é uma função pública")
}

// Função privada
func privada() {
    fmt.Println("Esta é uma função privada")
}
```

No exemplo acima, `Publica` pode ser acessada por outros pacotes, enquanto `privada` só pode ser usada dentro do pacote `exemplo`.

## 5. Pacotes Externos

Go permite o uso de pacotes externos para estender as funcionalidades do projeto. Para gerenciar essas dependências, utilizamos o `go get`.

### Instalando um Pacote Externo

Para baixar e instalar um pacote externo, utilize:

```sh
 go get [NOME_DO_PACOTE]
```

Por exemplo, para instalar o pacote `github.com/gorilla/mux`:

```sh
 go get github.com/gorilla/mux
```

Isso adicionará a dependência ao `go.mod` e baixará o código-fonte necessário para o diretório `go.sum`.

### Atualizando Pacotes

Para atualizar um pacote para a versão mais recente, utilize:

```sh
 go get -u [NOME_DO_PACOTE]
```

### Removendo Dependências Não Utilizadas

Após remover referências a um pacote do código, você pode limpar dependências não utilizadas com:

```sh
 go mod tidy
```

## 6. Variáveis

Go possui um sistema de tipagem estática, o que significa que cada variável deve ter um tipo definido. Existem diferentes formas de declarar variáveis na linguagem.

### Declaração Explícita

A declaração explícita de variáveis é feita utilizando a palavra-chave `var`, seguida do nome da variável, seu tipo e um valor opcional.

```go
var nome string = "GoLang"
var idade int = 25
```

Se um valor inicial for fornecido, o compilador já sabe o tipo da variável e a declaração pode ser feita sem o tipo explícito:

```go
var linguagem = "Go" // O tipo será inferido como string
```

### Inferência de Tipo

Go permite a inferência de tipo utilizando `:=`, que declara e inicializa a variável ao mesmo tempo, determinando seu tipo com base no valor atribuído.

```go
mensagem := "Olá, Go!" // O tipo inferido será string
numero := 42            // O tipo inferido será int
```

### Declaração Múltipla

É possível declarar múltiplas variáveis ao mesmo tempo:

```go
var (
    nome   string = "Alice"
    idade  int    = 30
    ativo  bool   = true
)
```

Ou utilizando inferência de tipo:

```go
nome, cidade := "Carlos", "São Paulo"
```

### Constantes

Constantes são declaradas com `const` e precisam ter seu valor definido no momento da declaração.

```go
const pi float64 = 3.14159
const mensagem = "Bem-vindo ao Go!"
```

### Troca de Valores

Go permite a troca de valores entre variáveis de maneira simples:

```go
a, b := "primeiro", "segundo"
a, b = b, a
fmt.Println(a, b) // "segundo primeiro"
```

## 7. Tipos de Dados

Go possui vários tipos de dados, que podem ser classificados em tipos numéricos, strings, booleanos e outros tipos compostos. Abaixo estão exemplos e explicações sobre cada um desses tipos.

### Tipos Numéricos

#### Números Inteiros

Go oferece diferentes tipos de inteiros com diferentes tamanhos e sinais:

- **int64**: Para números inteiros grandes, com sinal.
- **uint32**: Para números inteiros não negativos (sem sinal).
- **rune**: Alias para `int32`, usado para representar caracteres Unicode.
- **byte**: Alias para `uint8`, usado para representar bytes.

```go
var numero int64 = -100000000000
fmt.Println(numero)  // -100000000000

var numero2 uint32 = 10000
fmt.Println(numero2) // 10000

var numero3 rune = 123456
fmt.Println(numero3) // 123456

var numero4 byte = 123
fmt.Println(numero4) // 123
```

#### Números Reais

Go tem dois tipos para números de ponto flutuante:

- **float32**: Para números reais com precisão simples.
- **float64**: Para números reais com precisão dupla.

```go
var numeroReal1 float32 = 123.45
fmt.Println(numeroReal1) // 123.45

var numeroReal2 float64 = 1230000000.45
fmt.Println(numeroReal2) // 1230000000.45

numeroReal3 := 12345.67
fmt.Println(numeroReal3) // 12345.67
```

### Tipos de Texto

- **string**: Usado para representar texto, é uma sequência de caracteres Unicode.
- **rune**: Um `rune` é um valor do tipo `int32` e é usado para representar um único caractere Unicode.
- **byte**: Um `byte` é um `uint8` e é frequentemente usado para manipulação de dados binários.

```go
var str string = "Texto"
fmt.Println(str) // Texto

str2 := "Texto2"
fmt.Println(str2) // Texto2

char := 'B'
fmt.Println(char) // 66 (representação Unicode de 'B')
```

### Tipos Booleanos

O tipo **bool** é utilizado para valores verdadeiros ou falsos.

```go
var booleano1 bool
fmt.Println(booleano1) // false (valor padrão de bool)

booleano2 := true
fmt.Println(booleano2) // true
```

### Outros Tipos

- **error**: Go tem um tipo `error` incorporado, usado para representar erros.

```go
var erro error = errors.New("Erro interno")
fmt.Println(erro) // Erro interno
```

### Valores Não Inicializados

Quando uma variável é declarada, mas não inicializada com um valor, ela recebe um valor zero (valor padrão) para o seu tipo:

```go
var texto string
fmt.Println(texto) // (string vazia)

var booleano1 bool
fmt.Println(booleano1) // false
```
