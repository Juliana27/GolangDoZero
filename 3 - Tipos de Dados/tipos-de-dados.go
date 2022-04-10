package main

import (
	"errors"
	"fmt"
)

func main() {

	//int8, int16, int32, int64 a diferença é o tamanho do número que pode ser armazenado em cada um
	var numero int16 = 100
	fmt.Println(numero)

	//uint - unsygned int, int sem sinal, não aceita numeros negativos
	var numero2 uint32 = 10
	fmt.Println(numero2)

	//alias
	//int32 = rune
	var numero3 rune = 12345
	fmt.Println(numero3)

	//byte = uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	//float32, float64
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 12300000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12453.45
	fmt.Println(numeroReal3)

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	//no go não existe char
	char := 'B'
	fmt.Println(char)

	//toda variavel em com tem o seu valor iinicial o que é chamado de valor zero
	var texto string
	fmt.Println(texto)
	var num int16
	fmt.Println(num)

	//valor zero de bool é false
	var booleano bool = true
	fmt.Println(booleano)

	//valor zero é <nil>
	var erro error
	fmt.Println(erro)

	var erro2 error = errors.New("Erro interno")
	fmt.Println(erro2)
}
