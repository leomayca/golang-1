package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Leonardo"
	u.idade = 22
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos Bobos", 0}

	usuario2 := usuario{"Larissa", 22, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{idade: 40}
	fmt.Println(usuario3)
}