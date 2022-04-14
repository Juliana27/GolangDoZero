package main

import (
	"fmt"
	"time"
)

func main() {

	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando i ", i)
	// 	time.Sleep(time.Second) //"dormir" por 1 segundo
	// }

	// for j := 0; j < 10; j+=2 {
	// 	fmt.Println("Incrementando j ", j)
	// 	time.Sleep(time.Second)
	// }

	nomes := [3]string{"Juliana", "Vitor", "Thiago"}

	for indice, nome := range nomes { 	//posição e valor
		fmt.Println(indice, nome)
	}

	// for _, nome := range nomes { 	//não usando o indice
	// 	fmt.Println(nome)
	// }

	for indice, letra := range "PALAVRA"{
		fmt.Println(indice, letra) //assim vai trazer o hash das letras
	}

	for indice, letra := range "PALAVRA"{
		fmt.Println(indice, string(letra)) //assim vai trazer a letra mesmo
	}

	usuario := map[string]string {
		"nome": "Juliana",
		"sobrenome": "Brito",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	type usuarioStruct struct {
		nome string
		sobrenome string
	}

	// usuario2 := usuarioStruct{"Vitor", "Brito"}

	// for chave, valor := range usuario2 { //o Go não permite usar o for em struct
	// 	fmt.Println(chave, valor)
	// }

	for {
		fmt.Println("Loop Infinito")
		time.Sleep(time.Second)
	}
}