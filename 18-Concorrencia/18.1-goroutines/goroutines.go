package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Goroutines")

	// CONCORRÊNCIA != PARALELISMO

	go escrever("Hello World!") // goroutine
	escrever("Programando em Go!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
