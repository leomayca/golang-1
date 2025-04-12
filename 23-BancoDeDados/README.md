## 58. Conectando com o Banco de Dados

Go possui suporte nativo a SQL via o pacote `database/sql`, que permite conectar, consultar e manipular dados em bancos de dados relacionais, como MySQL e PostgreSQL.

Neste exemplo, veremos como conectar um programa Go a um banco de dados MySQL usando o driver `github.com/go-sql-driver/mysql`.

### Instalação do Driver

Para se conectar ao MySQL, você precisa instalar um driver compatível. Um dos mais utilizados é o:

```bash
go get github.com/go-sql-driver/mysql
```

### Estrutura básica de conexão

```go
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Formato: usuario:senha@/nomeBanco
	stringConexao := "usuario:senha@/nomedobanco"

	db, err := sql.Open("mysql", stringConexao)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verifica se a conexão foi bem-sucedida
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexão estabelecida com sucesso!")
}
```

### Realizando uma consulta simples

```go
linhas, err := db.Query("SELECT id, nome FROM usuarios")
if err != nil {
	log.Fatal(err)
}
defer linhas.Close()

for linhas.Next() {
	var id int
	var nome string

	if err := linhas.Scan(&id, &nome); err != nil {
		log.Fatal(err)
	}

	fmt.Println("ID:", id, "Nome:", nome)
}
```

#### Observações:

- `Query()` é usado para consultas que retornam múltiplas linhas (SELECT).
- `Scan()` mapeia cada coluna da linha atual para as variáveis fornecidas.
- `Next()` percorre cada linha do resultado.

### Inserindo dados

```go
resultado, err := db.Exec("INSERT INTO usuarios (nome) VALUES (?)", "João")
if err != nil {
	log.Fatal(err)
}

idInserido, _ := resultado.LastInsertId()
fmt.Println("Usuário inserido com ID:", idInserido)
```

### Atualizando e removendo

```go
_, err = db.Exec("UPDATE usuarios SET nome = ? WHERE id = ?", "Maria", 1)
if err != nil {
	log.Fatal(err)
}

_, err = db.Exec("DELETE FROM usuarios WHERE id = ?", 1)
if err != nil {
	log.Fatal(err)
}
```

### Considerações Finais

- Sempre feche a conexão com `defer db.Close()`.
- Use `?` como marcador de parâmetro para evitar SQL Injection.
- Verifique erros em cada operação para garantir a robustez da aplicação.
- `database/sql` é apenas uma interface, portanto é necessário utilizar um driver (como o `go-sql-driver/mysql`) para o banco específico.
