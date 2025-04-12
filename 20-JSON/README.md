## JSON

O Go oferece suporte nativo para codificação e decodificação de dados em formato JSON através do pacote `encoding/json`. Ele permite transformar **estruturas em JSON** (marshal) e também transformar **JSON em estruturas Go** (unmarshal).

### Convertendo uma struct para JSON (`json.Marshal`)

A função `json.Marshal` transforma uma estrutura (struct, mapa, slice, etc.) em um array de bytes no formato JSON.

#### Exemplo:

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Produto struct {
	Nome  string  `json:"nome"`
	Preco float64 `json:"preco"`
}

func main() {
	produto := Produto{"Notebook", 2499.90}

	jsonBytes, err := json.Marshal(produto)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonBytes))
}
```

#### Saída:

```
{"nome":"Notebook","preco":2499.9}
```

Note o uso da _tag_ `json:"nome"` para definir como o campo será representado no JSON.

### Convertendo JSON para struct (`json.Unmarshal`)

A função `json.Unmarshal` transforma um array de bytes contendo JSON em uma estrutura Go.

#### Exemplo:

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Produto struct {
	Nome  string  `json:"nome"`
	Preco float64 `json:"preco"`
}

func main() {
	jsonTexto := `{"nome":"Smartphone","preco":1299.99}`

	var p Produto

	err := json.Unmarshal([]byte(jsonTexto), &p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(p)
}
```

### Usando Mapas com JSON

Também é possível trabalhar com JSON usando mapas dinâmicos, especialmente quando não se conhece previamente a estrutura exata.

#### Marshal com mapa:

```go
m := map[string]interface{}{
	"nome":  "Câmera",
	"preco": 599.90,
}

jsonBytes, _ := json.Marshal(m)
fmt.Println(string(jsonBytes))
```

#### Unmarshal com mapa:

```go
jsonTexto := `{"nome":"Mouse","preco":89.90}`

var dados map[string]interface{}

err := json.Unmarshal([]byte(jsonTexto), &dados)
if err != nil {
	log.Fatal(err)
}

fmt.Println(dados["nome"])  // "Mouse"
fmt.Println(dados["preco"]) // 89.9
```

### Considerações importantes

- Os campos da struct precisam ser exportados (inicial com letra maiúscula) para que sejam acessíveis ao pacote `json`.
- A função `json.MarshalIndent` pode ser usada para gerar JSON formatado e mais legível.
- O tipo `interface{}` permite representar valores genéricos e pode ser útil quando o tipo do valor não é conhecido previamente.
