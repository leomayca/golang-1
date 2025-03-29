package main

import "fmt"

func main() {
	fmt.Println("Anônimas")

	func() {
		fmt.Println("Olá Mundo!")
	}()

	func(texto string) {
		fmt.Println(texto)
	}("Passando Parâmetro")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Parâmetro")

	fmt.Println(retorno)
}
