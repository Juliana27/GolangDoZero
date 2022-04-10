package main

import "fmt"

func main() {

	//ARITMÉTICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDivisao)

	// var numero1 int8 = 10
	// var numero2 int16 = 25
	// soma2 := numero1 + numero2 não se pode realizar operações matematicas com tipos diferentes de numeros

	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	//ATRIBUIÇÃO
	var variavel1 string = "String"
	variavel2 := "String 2"
	fmt.Println(variavel1, variavel2)

	//RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	//LÓGICOS
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)

	//UNÁRIOS
	numero := 10
	numero++
	numero += 15 // numero = numero + 15

	numero--
	numero -= 20 // numero = numero - 20
	numero *= 3  // numero = numero * 3
	numero /= 10
	numero %= 3
	fmt.Println(numero)

}
