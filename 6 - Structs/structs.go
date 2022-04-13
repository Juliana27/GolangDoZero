package main

import "fmt"

type usuario struct {
	nome     string
	idade    int8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     int8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Juliana"
	u.idade = 29
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua x", 1}

	usuario2 := usuario{"Juliana", 29, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Juliana"}
	fmt.Println(usuario3)
}
