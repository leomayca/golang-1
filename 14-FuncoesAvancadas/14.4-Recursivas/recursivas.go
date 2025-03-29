package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fmt.Println("Recursivas")

	// 1 1 2 3 5 8 13 21 34

	posicao := uint(11)
	fmt.Println(fibonacci(posicao))

	posicao2 := uint(7)

	for i := uint(1); i <= posicao2; i++ {
		fmt.Println(fibonacci(i))
	}
}
