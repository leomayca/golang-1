package main

import "fmt"

func soma(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	fmt.Println("Variáticas")

	totalDaSoma := soma(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(totalDaSoma)

	escrever("Olá Mundo", 10, 9, 8, 7, 6)
}
