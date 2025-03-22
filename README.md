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

Essas são as principais formas de declaração e uso de variáveis em Go, facilitando a organização e legibilidade do código.
