# Aplicação de Linha de Comando - Buscar IPs e Servidores

Este projeto é uma aplicação simples de linha de comando desenvolvida em Go. Ele permite buscar os **IPs** e **servidores** de um domínio fornecido na internet.

## Funcionalidades

A aplicação permite realizar as seguintes operações:

- **Buscar IPs**: Realiza a busca de IPs para um determinado nome de host.
- **Buscar Servidores**: Realiza a busca de servidores de nome (NS) para um determinado nome de host.

## Dependência

A aplicação utiliza a biblioteca `github.com/urfave/cli` para construção da linha de comando.

- Para instalar a dependência, basta rodar o comando:

```bash
go get github.com/urfave/cli
```

## Como Usar

1. Clone o repositório ou baixe os arquivos.
2. Execute o seguinte comando para rodar a aplicação com os parâmetros desejados:

```bash
go run main.go
```

### Comandos Disponíveis

- **Buscar IPs**:

```bash
go run main.go ip --host=devbook.com.br
```

Este comando vai buscar e exibir os IPs associados ao domínio fornecido (por padrão, será utilizado `devbook.com.br`).

- **Buscar Servidores**:

```bash
go run main.go servidores --host=devbook.com.br
```

Este comando vai buscar e exibir os servidores de nome (NS) associados ao domínio fornecido (por padrão, será utilizado `devbook.com.br`).

### Flags

A aplicação também possui uma flag `--host`, que pode ser usada para especificar o domínio que você deseja buscar os IPs ou servidores. Caso a flag não seja informada, o valor padrão será `devbook.com.br`.

- **Exemplo de uso da flag**:

```bash
go run main.go ip --host=example.com
```

## Estrutura do Código

- **main.go**: O ponto de entrada da aplicação. Ele configura a aplicação e executa com base nos argumentos de linha de comando.
- **app/app.go**: Contém a definição da aplicação CLI, com os comandos e flags.
- **funções**:
  - **buscarIps**: Função que busca os IPs de um domínio.
  - **buscarServidores**: Função que busca os servidores de nome (NS) de um domínio.

## Exemplo de Saída

### Busca de IPs:

```bash
$ go run main.go ip --host=google.com
172.217.14.238
172.217.14.238
...
```

### Busca de Servidores:

```bash
$ go run main.go servidores --host=google.com
ns1.google.com.
ns2.google.com.
...
```
