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
