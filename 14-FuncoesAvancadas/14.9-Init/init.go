package main

import "fmt"

var n int

func init() {
	fmt.Println("Executando a função init")
	n = 10
}

func main() {
	fmt.Println("Init")
	fmt.Println("Função main sendo executada")
	fmt.Println(n)
}
