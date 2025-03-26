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

## 8. Funções

As funções em Go são blocos de código que executam uma tarefa específica e podem retornar valores. Elas são definidas com a palavra-chave `func`, seguida pelo nome da função, parâmetros e tipo de retorno (se houver).

### Declaração de Função

Uma função em Go é declarada utilizando a palavra-chave `func`, seguida de seu nome, parâmetros entre parênteses e o tipo de retorno (se houver).

#### Exemplo de função simples:

```go
func somar(n1 int, n2 int) int {
    return n1 + n2
}
```

Aqui, a função `somar` recebe dois parâmetros (`n1` e `n2`), ambos do tipo `int`, e retorna um valor do tipo `int`.

### Função com múltiplos retornos

Go permite que funções retornem múltiplos valores. Isso é útil quando se deseja retornar mais de um resultado de uma função.

#### Exemplo de função com múltiplos retornos:

```go
func calcularOperacoes(n1, n2 int) (int, int, int) {
    soma := n1 + n2
    subtracao := n1 - n2
    multiplicacao := n1 * n2
    return soma, subtracao, multiplicacao
}
```

Neste exemplo, a função `calcularOperacoes` retorna três valores: a soma, a subtração e a multiplicação de dois números.

### Funções Anônimas

Go também permite a criação de funções anônimas (sem nome), que podem ser atribuídas a variáveis e chamadas posteriormente.

#### Exemplo de função anônima:

```go
var saudacao = func(nome string) string {
    return "Olá, " + nome + "!"
}
```

Aqui, a função anônima recebe um parâmetro `nome` do tipo `string`, cria uma saudação e retorna essa string.

### Chamada de Funções

Após definir as funções, você pode chamá-las no corpo do programa.

#### Exemplo de chamada de função:

```go
func main() {
    // Chama a função somar
    resultadoSoma := somar(5, 7)
    fmt.Println("Resultado da soma:", resultadoSoma)

    // Chama a função anônima
    resultadoSaudacao := saudacao("Carlos")
    fmt.Println(resultadoSaudacao)  // "Olá, Carlos!"

    // Chama a função com múltiplos retornos
    soma, subtracao, multiplicacao := calcularOperacoes(10, 3)
    fmt.Println("Soma:", soma, "Subtração:", subtracao, "Multiplicação:", multiplicacao)
}
```

### Descartando valores não necessários

No caso de funções que retornam múltiplos valores, você pode ignorar um dos retornos utilizando o caracter `_`. No exemplo abaixo, apenas o valor da soma é utilizado, e os outros dois valores são descartados:

```go
resultadoSoma, _, _ := calcularOperacoes(10, 3)
fmt.Println("Resultado da soma:", resultadoSoma)  // Apenas a soma será exibida
```

## 9. Operadores

Os operadores em Go são usados para realizar operações sobre variáveis e valores. Eles podem ser divididos em várias categorias, como operadores aritméticos, de atribuição, relacionais, lógicos, unários e ternários. Abaixo estão os principais operadores e exemplos de como usá-los de forma clara.

### Operadores Aritméticos

São usados para realizar operações matemáticas básicas, como soma, subtração, multiplicação, divisão e o cálculo do resto da divisão.

#### Exemplos de operadores aritméticos:

```go
// Soma de dois números
soma := 10 + 5

// Subtração de dois números
subtracao := 15 - 4

// Multiplicação de dois números
multiplicacao := 7 * 3

// Divisão de dois números
divisao := 20 / 4

// Resto da divisão
resto := 9 % 4

fmt.Println("Soma:", soma)
fmt.Println("Subtração:", subtracao)
fmt.Println("Multiplicação:", multiplicacao)
fmt.Println("Divisão:", divisao)
fmt.Println("Resto:", resto)
```

### Operadores de Atribuição

Os operadores de atribuição são usados para dar valores a variáveis. Além da atribuição simples, Go oferece operadores de atribuição compostos.

#### Exemplos de operadores de atribuição:

```go
// Atribuição simples
var x int = 10

// Atribuição com incremento
x += 5 // x = x + 5

// Atribuição com multiplicação
x *= 2 // x = x * 2

fmt.Println("Valor de x:", x)
```

### Operadores Relacionais

Esses operadores são usados para comparar dois valores e retornar um valor booleano (`true` ou `false`), dependendo do resultado da comparação.

#### Exemplos de operadores relacionais:

```go
// Verificando se os números são iguais
igual := 10 == 10 // true

// Verificando se os números são diferentes
diferente := 10 != 5 // true

// Verificando se um número é maior que o outro
maior := 10 > 5 // true

// Verificando se um número é menor ou igual ao outro
menorOuIgual := 10 <= 15 // true

fmt.Println("São iguais?", igual)
fmt.Println("São diferentes?", diferente)
fmt.Println("É maior?", maior)
fmt.Println("É menor ou igual?", menorOuIgual)
```

### Operadores Lógicos

Os operadores lógicos são usados para combinar expressões booleanas. Eles permitem realizar operações **E** (`&&`), **OU** (`||`) e **NÃO** (`!`).

#### Exemplos de operadores lógicos:

```go
verdadeiro, falso := true, false

// Operador E (AND) - verdadeiro se ambos os lados forem verdadeiros
e := verdadeiro && falso // false

// Operador OU (OR) - verdadeiro se ao menos um lado for verdadeiro
ou := verdadeiro || falso // true

// Operador NÃO (NOT) - inverte o valor lógico
nao := !verdadeiro // false

fmt.Println("Resultado E:", e)
fmt.Println("Resultado OU:", ou)
fmt.Println("Resultado NÃO:", nao)
```

### Operadores Unários

Os operadores unários são usados para realizar operações em um único operando, como incremento (`++`), decremento (`--`) e operações combinadas de atribuição.

#### Exemplos de operadores unários:

```go
numero := 5

// Incremento
numero++ // numero = numero + 1

// Decremento
numero-- // numero = numero - 1

// Operação de atribuição combinada
numero += 10 // numero = numero + 10

fmt.Println("Valor final de numero:", numero)
```

### Operador Ternário (Simulado)

Embora Go não tenha um operador ternário como em outras linguagens, podemos simular seu comportamento com uma estrutura condicional simples.

#### Exemplo de operador ternário simulado:

```go
numero := 8
var resultado string

if numero > 5 {
    resultado = "Maior que 5"
} else {
    resultado = "Menor ou igual a 5"
}

fmt.Println("Resultado:", resultado)
```

## 10. Structs

Structs são tipos compostos em Go, usados para agrupar dados diferentes em uma única estrutura. Cada elemento de uma struct é chamado de **campo** e pode ter tipos diferentes. Structs são muito úteis para modelar entidades mais complexas, como usuários, produtos ou qualquer outra coisa que precise agrupar informações relacionadas.

### Definindo uma Struct

Para definir uma struct, você usa a palavra-chave `type`, seguida pelo nome da struct, o tipo de dados de cada campo e os campos que ela terá.

#### Exemplo de definição de struct:

```go
type Usuario struct {
    Nome    string
    Idade   uint8
    Endereco Endereco
}

type Endereco struct {
    Logradouro string
    Numero     uint8
}
```

### Criando e Inicializando Structs

Uma struct pode ser inicializada de várias maneiras. Você pode usar o nome do tipo e definir cada campo individualmente, ou usar uma inicialização direta com valores.

#### Exemplo 1: Inicialização com atribuição direta dos campos

```go
var u Usuario
u.Nome = "Leonardo"
u.Idade = 22
fmt.Println(u)
```

Neste exemplo, criamos uma variável `u` do tipo `Usuario` e atribuimos valores a cada campo da struct.

#### Exemplo 2: Inicialização com valores ao criar a struct

```go
enderecoExemplo := Endereco{"Rua dos Bobos", 0}
usuario2 := Usuario{"Larissa", 22, enderecoExemplo}
fmt.Println(usuario2)
```

Aqui, inicializamos a struct `Endereco` e passamos os valores diretamente ao criar a struct `Usuario`.

#### Exemplo 3: Inicialização parcial de uma struct

Go permite inicializar uma struct, especificando apenas os campos que você deseja atribuir. Os campos não especificados recebem valores zero (como `0` para `int`, `""` para `string`, etc).

```go
usuario3 := Usuario{Idade: 40}
fmt.Println(usuario3)
```

Neste caso, a struct `Usuario` foi criada e o campo `Idade` foi atribuído com o valor `40`, enquanto o campo `Nome` e o campo `Endereco` terão seus valores padrão (em branco e `Endereco{}`).

### Acessando Campos de uma Struct

Após a criação de uma struct, você pode acessar seus campos diretamente usando o operador `.`.

#### Exemplo de acesso a campos:

```go
fmt.Println("Nome do usuário:", usuario2.Nome)
fmt.Println("Endereço do usuário:", usuario2.Endereco.Logradouro)
```

## 11. Herança (Composição de Structs)

Diferente de linguagens orientadas a objetos que possuem herança de classes, Go utiliza **composição** para reaproveitar código e estruturar dados de forma hierárquica.

Em Go, uma struct pode conter outra struct como um campo, permitindo o compartilhamento de atributos sem precisar de herança tradicional.

### Definição de Structs com Composição

Podemos criar uma struct `Pessoa` que contenha informações comuns, como nome e idade, e depois utilizá-la dentro de uma struct `Estudante`, que adiciona mais informações específicas.

#### Exemplo:

```go
type Pessoa struct {
    Nome      string
    Sobrenome string
    Idade     uint8
    Altura    uint8
}

type Estudante struct {
    Pessoa    // Composição: reaproveitando os campos da struct Pessoa
    Curso     string
    Faculdade string
}
```

### Criando e Utilizando Structs Compostas

```go
func main() {
    pessoa1 := Pessoa{"João", "Silva", 20, 178}
    fmt.Println(pessoa1)

    estudante1 := Estudante{Pessoa: pessoa1, Curso: "Engenharia", Faculdade: "USP"}
    fmt.Println(estudante1)

    // Podemos acessar os campos da struct interna diretamente
    fmt.Println("Nome do estudante:", estudante1.Nome)
    fmt.Println("Curso:", estudante1.Curso)
}
```

### Como Funciona?

- A struct `Estudante` **incorpora** a struct `Pessoa`, permitindo acesso direto aos seus campos sem precisar de uma sintaxe adicional.
- O acesso aos campos da struct embutida (`Pessoa`) pode ser feito diretamente como se fossem da struct `Estudante`.
- Essa abordagem permite reutilizar código sem precisar de herança tradicional.

### Diferença Entre Composição e Herança

| Característica  | Herança (em outras linguagens)                                  | Composição (Go)                                                     |
| --------------- | --------------------------------------------------------------- | ------------------------------------------------------------------- |
| Reuso de código | Classes herdam atributos e métodos de uma superclasse           | Structs incorporam outras struct sem hierarquia fixa                |
| Flexibilidade   | Relacionamento fixo (classe filha depende da classe pai)        | Relacionamento flexível (pode combinar struct de formas diferentes) |
| Encapsulamento  | Possui visibilidade entre classes (pública, protegida, privada) | Go usa exportação por convenção (letra maiúscula para público)      |

## 12. Arrays e Slices

Em Go, **arrays** e **slices** são estruturas usadas para armazenar coleções de elementos do mesmo tipo. Embora sejam semelhantes, possuem diferenças importantes em relação ao tamanho e flexibilidade.

### Arrays

Um **array** em Go é uma coleção de elementos de tamanho fixo. O tamanho do array faz parte do seu tipo, ou seja, arrays de tamanhos diferentes são considerados tipos distintos.

#### Declarando Arrays

```go
var numeros [5]int  // Array de inteiros com 5 posições
numeros[0] = 10     // Atribuindo um valor
fmt.Println(numeros) // [10 0 0 0 0]
```

Também podemos inicializar um array diretamente com valores:

```go
nomes := [3]string{"Alice", "Bob", "Charlie"}
fmt.Println(nomes) // ["Alice", "Bob", "Charlie"]
```

Se quisermos que o tamanho do array seja inferido automaticamente:

```go
valores := [...]int{1, 2, 3, 4, 5}
fmt.Println(valores) // [1 2 3 4 5]
```

### Slices

Um **slice** é semelhante a um array, mas com tamanho dinâmico. Ele é baseado em um array subjacente, permitindo redimensionamento e operações mais flexíveis.

#### Criando um Slice

```go
numeros := []int{10, 20, 30, 40, 50}
fmt.Println(numeros) // [10 20 30 40 50]
```

Note que não especificamos o tamanho do slice (`[]int`), diferentemente do array (`[5]int`).

#### Adicionando Elementos ao Slice

Podemos usar `append` para adicionar novos elementos a um slice:

```go
numeros = append(numeros, 60, 70)
fmt.Println(numeros) // [10 20 30 40 50 60 70]
```

#### Criando um Slice a Partir de um Array

Podemos criar um slice extraindo parte de um array:

```go
array := [5]int{1, 2, 3, 4, 5}
slice := array[1:4] // Do índice 1 ao 3 (não inclui o 4)
fmt.Println(slice)  // [2 3 4]
```

Se alterarmos um elemento do array original, o slice também será afetado:

```go
array[2] = 99
fmt.Println(slice)  // [2 99 4]
```

Isso ocorre porque slices compartilham a mesma área de memória do array original.

#### Verificando Tipos

Podemos usar o pacote `reflect` para verificar se uma variável é um array ou slice:

```go
import "reflect"

fmt.Println(reflect.TypeOf(array)) // [5]int
fmt.Println(reflect.TypeOf(slice)) // []int
```

### Diferenças Entre Arrays e Slices

| Característica       | Arrays                              | Slices                               |
| -------------------- | ----------------------------------- | ------------------------------------ |
| Tamanho              | Fixo                                | Dinâmico                             |
| Flexibilidade        | Menos flexível                      | Mais flexível                        |
| Compartilha memória? | Não                                 | Sim (se criado a partir de um array) |
| Uso recomendado      | Quando o tamanho é conhecido e fixo | Quando o tamanho pode mudar          |

## 13. Ponteiros

Em Go, um **ponteiro** é uma variável que armazena o endereço de memória de outra variável. Isso permite modificar valores diretamente na memória, evitando cópias desnecessárias de dados.

### Declarando Ponteiros

Para declarar um ponteiro, utilizamos o operador `*` antes do tipo da variável:

```go
var ponteiro *int  // Ponteiro para um valor do tipo int
```

O valor inicial de um ponteiro sem atribuição é `nil` (nulo), pois ele ainda não aponta para nenhum endereço válido.

### Obtendo o Endereço de Memória

Podemos obter o endereço de uma variável utilizando o operador `&`:

```go
numero := 10
ponteiro := &numero // Ponteiro recebe o endereço da variável numero
fmt.Println("Endereço de memória de numero:", ponteiro)
```

### Desreferenciando um Ponteiro

A desreferenciação (`*ponteiro`) nos permite acessar o valor armazenado no endereço apontado pelo ponteiro:

```go
fmt.Println("Valor armazenado no ponteiro:", *ponteiro) // Saída: 10
```

### Modificando Valores Usando Ponteiros

Quando alteramos o valor da variável original, o ponteiro reflete essa mudança:

```go
numero = 20
fmt.Println("Valor atualizado:", *ponteiro) // Saída: 20
```

Também podemos modificar o valor da variável original através do ponteiro:

```go
*ponteiro = 30
fmt.Println("Novo valor de numero:", numero) // Saída: 30
```

### Exemplo Completo:

```go
package main

import "fmt"

func main() {
    var numero int = 100
    var ponteiro *int

    fmt.Println("Valor de numero:", numero)
    fmt.Println("Valor do ponteiro (antes da atribuição):", ponteiro)

    ponteiro = &numero // Ponteiro recebe o endereço da variável numero
    fmt.Println("Endereço de numero:", ponteiro)
    fmt.Println("Valor apontado pelo ponteiro:", *ponteiro)

    *ponteiro = 200 // Modificando o valor diretamente na memória
    fmt.Println("Novo valor de numero:", numero)
}
```

### Quando Usar Ponteiros?

- Para **evitar cópias desnecessárias** de valores grandes na memória.
- Para **modificar valores dentro de funções**, pois em Go os argumentos são passados por cópia por padrão.
- Para **trabalhar com estruturas de dados mais eficientes**, como listas e árvores.

### Diferença Entre Valor e Referência

| Tipo de Atribuição | Comportamento                                                                                                  |
| ------------------ | -------------------------------------------------------------------------------------------------------------- |
| Atribuição Direta  | Cria uma cópia do valor original. Modificações na cópia não afetam o valor original.                           |
| Uso de Ponteiros   | O ponteiro armazena o endereço do valor original. Modificações feitas via ponteiro afetam a variável original. |
