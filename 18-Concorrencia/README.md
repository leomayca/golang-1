## 34. Introdução às Goroutines

Goroutines são o coração da concorrência em Go, permitindo que você execute várias tarefas ao mesmo tempo de forma simples e eficiente. Imagine que você está cozinhando: enquanto a água ferve em uma panela, você corta os legumes em outra tarefa — as goroutines são como esses "ajudantes" que trabalham juntos, mas de forma independente, sem precisar de threads pesadas como em outras linguagens.

Em Go, uma goroutine é uma função que roda concorrentemente, gerenciada pela runtime da linguagem. Ela é leve (usa poucos kilobytes de memória) e o escalonador de Go decide quando cada uma executa, tornando a concorrência prática e acessível.

### Exemplo Didático: Contando e Cantando

Vamos ver um exemplo básico para entender como as goroutines funcionam:

```go
package main

import (
	"fmt"
	"time"
)

func contar() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Contando:", i)
		time.Sleep(500 * time.Millisecond) // Pausa de meio segundo
	}
}

func cantar() {
	for i := 1; i <= 3; i++ {
		fmt.Println("Cantando: Lá lá lá", i)
		time.Sleep(800 * time.Millisecond) // Pausa de 0,8 segundo
	}
}

func main() {
	fmt.Println("Início do programa!")

	go contar()  // Inicia a contagem em uma goroutine
	cantar()    // Canta no thread principal

	fmt.Println("Fim do programa!")
}
```

### O Que Acontece Aqui?

1. **`go contar()`:**

   - A função `contar` roda em uma goroutine, ou seja, começa a contar de 1 a 5 em "segundo plano". Cada número é impresso com uma pausa de 0,5 segundo.

2. **`cantar()`:**

   - Executada no thread principal (sem `go`), essa função imprime "Lá lá lá" três vezes, com pausas de 0,8 segundo.

3. **Saída Esperada:**
   - Você verá algo como:
     ```
     Início do programa!
     Contando: 1
     Cantando: Lá lá lá 1
     Contando: 2
     Contando: 3
     Cantando: Lá lá lá 2
     Contando: 4
     Contando: 5
     Cantando: Lá lá lá 3
     Fim do programa!
     ```
   - As mensagens se intercalam porque as duas funções rodam concorrentemente, mas o escalonador de Go decide a ordem exata.

### Por Que Isso É Útil?

Imagine que `contar` é uma tarefa demorada, como baixar dados da internet, e `cantar` é uma interface que precisa responder ao usuário. Com goroutines, você pode fazer ambos ao mesmo tempo sem travar o programa!

### Atenção

Nesse exemplo, o programa pode acabar antes da goroutine `contar` terminar, porque o `main` não espera por ela. Para evitar isso, podemos adicionar sincronização (como um `time.Sleep` ou um `WaitGroup`), mas isso será explorado em tópicos futuros.

## 35. Wait Group

Em Go, **goroutines** são executadas de forma assíncrona. No entanto, isso pode causar situações em que o programa principal termine antes que as goroutines finalizem suas execuções. Para resolver esse problema, podemos utilizar o **WaitGroup**, que permite aguardar a finalização de múltiplas goroutines antes de continuar a execução do programa.

### O que é `sync.WaitGroup`?

O `WaitGroup` é uma estrutura da biblioteca `sync` que permite controlar e aguardar a finalização de goroutines. Com ele, é possível adicionar contadores para cada tarefa em paralelo e esperar até que todas elas sejam concluídas.

### Exemplo Básico

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1) // Adiciona uma tarefa ao contador

	go func() {
		fmt.Println("Executando tarefa em goroutine")
		wg.Done() // Informa que a tarefa terminou
	}()

	wg.Wait() // Aguarda até que todas as tarefas terminem
	fmt.Println("Todas as goroutines finalizaram")
}
```

### Exemplo com Múltiplas Goroutines e Espera

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	tarefas := []string{
		"Tarefa 1 - Processar dados",
		"Tarefa 2 - Gerar relatório",
		"Tarefa 3 - Enviar email",
	}

	for _, tarefa := range tarefas {
		wg.Add(1)
		go func(t string) {
			defer wg.Done()
			fmt.Println(t, "iniciada")
			time.Sleep(2 * time.Second)
			fmt.Println(t, "finalizada")
		}(tarefa)
	}

	wg.Wait()
	fmt.Println("Todas as tarefas foram concluídas")
}
```

### Explicação dos métodos:

- `wg.Add(n)`: Informa que n goroutines serão aguardadas.
- `wg.Done()`: Deve ser chamado ao final de cada goroutine para sinalizar que terminou.
- `wg.Wait()`: Faz o programa esperar até que todas as goroutines finalizem (`Done()` seja chamado para cada `Add`).

### Dica Importante

Sempre que usar `WaitGroup` com goroutines dentro de loops, **lembre-se de passar variáveis como parâmetro**, como feito acima (`(t string)`), para evitar o problema de _variável compartilhada_ (closure com valor inesperado).

### Quando usar `WaitGroup`?

- Quando você precisa aguardar que várias tarefas assíncronas finalizem.
- Quando está trabalhando com processamento concorrente que precisa ser sincronizado no final.
- Em cenários com paralelismo, mas onde a ordem de execução não importa, apenas a conclusão completa.

## 36. Canais

**Canais (Channels)** em Go são usados para permitir a **comunicação segura entre goroutines**. Eles servem como um meio para que uma goroutine envie dados e outra goroutine os receba, garantindo sincronização sem a necessidade de mecanismos manuais de bloqueio.

### Criando um Canal

Para criar um canal, usamos a função `make`:

```go
canal := make(chan string)
```

Isso cria um canal capaz de transportar valores do tipo `string`.

### Enviando e Recebendo dados

- **Enviar**: `canal <- "mensagem"`
- **Receber**: `msg := <-canal`

### Exemplo Básico com Goroutine

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Golang é top!", canal)

	for msg := range canal {
		fmt.Println(msg)
	}

	fmt.Println("Fim do programa")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 3; i++ {
		canal <- fmt.Sprintf("%d - %s", i+1, texto)
		time.Sleep(time.Second)
	}
	close(canal) // Fecha o canal após o envio
}
```

### Explicação:

- A função `escrever` envia uma mensagem para o canal a cada segundo.
- A função `main` escuta o canal com um `for range`, imprimindo cada mensagem recebida.
- O canal é fechado com `close(canal)` para indicar que nenhum outro dado será enviado. Isso evita deadlocks e permite que o `range` pare de escutar.

### Forma Alternativa: Recebendo com Verificação

```go
for {
	mensagem, aberto := <-canal
	if !aberto {
		break
	}
	fmt.Println(mensagem)
}
```

Essa forma permite checar se o canal ainda está aberto.

### Comunicação entre Goroutines

Canais são úteis para sincronizar tarefas:

```go
done := make(chan bool)

go func() {
	// fazer algo
	time.Sleep(2 * time.Second)
	done <- true
}()

<-done // Aguarda até que a goroutine envie sinal
```

### Quando Usar Canais?

- Para sincronizar tarefas assíncronas.
- Quando uma goroutine precisa esperar dados de outra.
- Para evitar o uso manual de mutexes ou waitgroups em fluxos simples de comunicação.

## 37. Canais com Buffer

Canais com buffer (ou canais **bufferizados**) em Go permitem enviar valores mesmo que **nenhuma goroutine esteja imediatamente pronta para recebê-los**. Eles armazenam os dados até que sejam recebidos, funcionando como uma fila temporária.

### Criando um Canal com Buffer

Você define o tamanho do buffer ao criar o canal com `make`:

```go
canal := make(chan string, 2)
```

Nesse exemplo, o canal pode armazenar **até 2 valores** antes que qualquer recebedor seja necessário.

### Envio e Recebimento

Com canais bufferizados, você pode fazer:

```go
canal <- "Mensagem 1"
canal <- "Mensagem 2"
```

Mesmo sem nenhuma goroutine recebendo imediatamente. Os valores ficam armazenados no canal até serem lidos.

### Exemplo Prático

```go
package main

import "fmt"

func main() {
	fmt.Println("Canais com buffer")

	canal := make(chan string, 2)

	canal <- "Hello World!"
	canal <- "Programando em Go!"

	mensagem1 := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem1)
	fmt.Println(mensagem2)
}
```

**Saída:**

```
Canais com buffer
Hello World!
Programando em Go!
```

### O Que Acontece se Ultrapassar o Buffer?

Se você tentar enviar mais dados do que o canal suporta **sem receber nada**, o programa ficará bloqueado até que um dado seja retirado do canal:

```go
canal := make(chan int, 1)

canal <- 1
canal <- 2 // Aqui o programa trava até alguém ler o primeiro valor
```

### Usando com Goroutines

Canais com buffer também podem ser usados para comunicação entre goroutines com mais flexibilidade:

```go
func escrever(canal chan string) {
	canal <- "Escrevendo 1"
	canal <- "Escrevendo 2"
	canal <- "Escrevendo 3"
	close(canal)
}

func main() {
	c := make(chan string, 3)
	go escrever(c)

	for msg := range c {
		fmt.Println(msg)
	}
}
```
