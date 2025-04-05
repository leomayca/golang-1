package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Canais")

	canal := make(chan string)
	go escrever("Hello World!", canal)

	// for {
	// 	mensagem, aberto := <-canal
	// 	if !aberto {
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }

	// Faz a mesma coisa do for acima
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa!")
}

func escrever(texto string, canal chan string) {
	for range 5 {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal)
}
