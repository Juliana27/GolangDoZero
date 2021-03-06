package main

import "fmt"

func main() {

	var variavel1 string = "Variável 1"
	variavel2 := "Variável 2"

	var (
		variavel3 string = "Variavel 3"
		variavel4 string = "Variavel 4"
	)

	variavel5, variavel6 := "Variavel 5", "Variavel 6"

	fmt.Println(variavel1)
	fmt.Println(variavel2)
	fmt.Println(variavel3, variavel4)
	fmt.Println(variavel5, variavel6)

	const constante1 string = "constante 1"
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
