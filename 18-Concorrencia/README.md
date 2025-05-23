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

## 38. Select

O `select` em Go é uma poderosa construção que permite **aguardar múltiplas operações de canal ao mesmo tempo**. Ele funciona como um `switch`, mas para canais. Assim que um dos canais estiver pronto para envio ou recebimento, o `select` executa o bloco de código correspondente.

### Exemplo Básico

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Select")

	canal1 := make(chan string)
	canal2 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Mensagem do Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Mensagem do Canal 2"
		}
	}()

	for {
		select {
		case msg1 := <-canal1:
			fmt.Println("Recebido:", msg1)
		case msg2 := <-canal2:
			fmt.Println("Recebido:", msg2)
		}
	}
}
```

### Como funciona?

- O `select` escuta os dois canais.
- Ele **executa o primeiro caso disponível**, ou seja, o primeiro canal que tiver uma mensagem pronta.
- Se **mais de um canal estiver pronto**, ele escolhe **aleatoriamente entre eles**.
- Se **nenhum canal estiver pronto**, o `select` **bloqueia até que algum esteja**.

### Exemplo com `default`

Você pode usar o `default` para evitar o bloqueio quando nenhum canal estiver pronto:

```go
select {
case msg := <-canal1:
	fmt.Println("Canal 1:", msg)
case msg := <-canal2:
	fmt.Println("Canal 2:", msg)
default:
	fmt.Println("Nenhum canal pronto, continuando execução...")
}
```

### Usos comuns do `select`:

- **Multiplexação** de canais (como no exemplo acima).
- **Timeouts** com `time.After`:

  ```go
  select {
  case msg := <-canal:
  	fmt.Println("Recebido:", msg)
  case <-time.After(time.Second * 3):
  	fmt.Println("Timeout após 3 segundos")
  }
  ```

- **Finalização de goroutines** usando canais de controle.

### Quando usar?

- Quando há múltiplas goroutines produzindo dados em canais diferentes.
- Quando você precisa lidar com canais em diferentes ritmos.
- Para evitar travamentos em canais que podem demorar a responder.

## 39. Padrões de Concorrência - Worker Pools

Em Go, **Worker Pools** são um padrão comum para otimizar o uso de goroutines e processar tarefas concorrentes de forma eficiente. Eles ajudam a **controlar a quantidade de trabalho sendo executado simultaneamente**, evitando o uso excessivo de recursos do sistema.

### O que é um Worker Pool?

Um _Worker Pool_ consiste em:

- Um **conjunto de workers** (goroutines) que processam tarefas.
- Um **canal de entrada** com as tarefas.
- Um **canal de saída** com os resultados processados.

Em vez de criar uma goroutine para cada tarefa (o que pode causar sobrecarga), você cria um número fixo de workers que pegam tarefas do canal e as executam.

### Exemplo de Worker Pool

```go
package main

import "fmt"

func main() {
	fmt.Println("Worker Pools")

	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	// Criando 3 workers
	for i := 0; i < 3; i++ {
		go worker(tarefas, resultados)
	}

	// Enviando tarefas (cálculo de fibonacci)
	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	// Recebendo os resultados
	for i := 0; i < 45; i++ {
		fmt.Println(<-resultados)
	}
}

func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
```

### O que está acontecendo aqui?

1. **Criamos os canais** `tarefas` e `resultados`.
2. **Três workers** são iniciados como goroutines com a função `worker`.
3. **45 tarefas** (números inteiros) são enviadas para o canal `tarefas`.
4. Cada worker consome valores do canal `tarefas` e envia o resultado (o valor de Fibonacci) para o canal `resultados`.
5. O programa imprime os 45 resultados.

### Vantagens do Worker Pool

- Evita criar goroutines demais.
- Controla melhor o consumo de recursos.
- Facilita o balanceamento de carga entre workers.
- É mais previsível e escalável em sistemas concorrentes.

### Quando usar?

- Quando você tem **muitas tarefas independentes** a serem executadas.
- Quando quer **limitar o número de goroutines** que rodam simultaneamente.
- Ao processar filas de mensagens, jobs em segundo plano, ou requisições em lote.

### Dica

Você pode ajustar o número de workers (`go worker(...)`) de acordo com a carga esperada e a capacidade do seu sistema.

## 40. Padrões de Concorrência - Generator

O padrão de concorrência **Generator** em Go permite que funções gerem uma sequência de valores sob demanda, utilizando **canais** e **goroutines**.

Esse padrão é útil quando queremos criar um fluxo contínuo ou preguiçoso (_lazy_) de dados sem bloquear a execução principal.

### O que é um Generator?

Um _Generator_ é uma função que **retorna um canal de saída** (`<-chan tipo`) e envia valores continuamente (ou conforme a lógica), geralmente dentro de uma goroutine.

### Exemplo Prático

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Generators")

	canal := escrever("Hello World!")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}
```

### Explicação

- A função `escrever` **retorna um canal** por onde as mensagens serão enviadas.
- Dentro da função anônima (goroutine), o texto é enviado repetidamente com um pequeno intervalo.
- No `main`, o canal é consumido em um loop, pegando os 10 primeiros valores.

### Características do Generator

- Usa **goroutines** para executar a geração em segundo plano.
- Usa **canais** para comunicar valores gerados.
- É útil quando você quer desacoplar a **produção de dados** do seu **consumo**.
- Ideal para criar _pipelines de processamento_.

### Quando usar?

- Para gerar sequências infinitas ou semi-infinita de valores.
- Para transformar uma lógica de produção de dados em um componente reutilizável.
- Em sistemas reativos, pipelines ou simulações.

### Exemplo com Contador

Um gerador que produz números incrementais:

```go
func contador() <-chan int {
	canal := make(chan int)

	go func() {
		for i := 0; ; i++ {
			canal <- i
			time.Sleep(time.Millisecond * 300)
		}
	}()

	return canal
}

func main() {
	numeros := contador()
	for i := 0; i < 5; i++ {
		fmt.Println(<-numeros)
	}
}
```

## 41. Padrões de Concorrência - Multiplexador

O padrão de concorrência **Multiplexador** em Go tem como objetivo **combinar múltiplos canais de entrada** em **um único canal de saída**. Ele permite que valores oriundos de diferentes origens sejam processados de forma centralizada, sem perder a concorrência entre as fontes.

### Quando usar um Multiplexador?

- Quando você quer **consolidar** dados de múltiplas goroutines.
- Quando precisa **observar várias fontes de dados simultaneamente**.
- Para simplificar o consumo de eventos concorrentes em um único ponto.

### Exemplo Prático

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Multiplexador")

	canal := multiplexar(escrever("Hello World!"), escrever("Programando em Go!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}

func multiplexar(canal1, canal2 <-chan string) <-chan string {
	saida := make(chan string)

	go func() {
		for {
			select {
			case msg := <-canal1:
				saida <- msg
			case msg := <-canal2:
				saida <- msg
			}
		}
	}()

	return saída
}
```

### Explicação

- A função `escrever` retorna um canal que emite mensagens com delays aleatórios.
- A função `multiplexar` recebe dois canais como entrada e retorna **um canal único de saída**.
- Dentro de uma goroutine, o `select` escuta os dois canais e envia qualquer mensagem recebida para o canal final.
- No `main`, o programa consome 10 mensagens do canal único, que vêm intercaladas das duas fontes.

### Visualizando o Fluxo

```
escrever("Hello World!") ┐
                          ├──> multiplexar() ──> canal final ──> consumidor
escrever("Programando!") ┘
```

### Variação com Múltiplos Canais

Você pode generalizar o padrão usando _slices_ de canais e reflet, mas a versão acima cobre bem os conceitos principais para a maioria dos usos.
