package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"Juliana", "Brito", 29, 157}
	fmt.Println(p1)

	e1 := estudante{p1, "TADS", "Estácio"}
	fmt.Println(e1.nome)
	fmt.Println(e1.altura)
	fmt.Println(e1)
}
