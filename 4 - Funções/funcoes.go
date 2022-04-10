package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

//mais de um retorno na função
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("texto da função")
	fmt.Println(resultado)

	//quando chamar uma função que tem mais de um retorno, porém você não quer todos os retornos, usar _
	// resultadoSoma, _ := calculosMatematicos(10, 15)
	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)
}
