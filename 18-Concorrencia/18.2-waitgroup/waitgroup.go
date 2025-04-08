package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Waitgroup")

	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		escrever("Hello World!")
		waitGroup.Done()
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func escrever(texto string) {
	for range 5 {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
