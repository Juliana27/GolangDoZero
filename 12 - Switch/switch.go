package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sabado"
	default:
		return "Número inválido"
	}

}

func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segnda-feira"
	case numero == 3:
		return "Terça-feira"
	case numero == 4:
		return "Quarte-feira"
	case numero == 5:
		return "Quinta-feira"
	case numero == 6:
		return "Sexta-feira"
	case numero == 7:
		return "Sabado"
	default:
		return "Número inválido"
	}
}

	func diaDaSemana3(numero int) string {

		var diaDaSemana string
		
		switch {
		case numero == 1:
			diaDaSemana = "Domingo"
			fallthrough		//faz seu código cair no próximo caso do switch
		case numero == 2:
			diaDaSemana = "Segnda-feira"
		case numero == 3:
			diaDaSemana = "Terça-feira"
		case numero == 4:
			diaDaSemana = "Quarte-feira"
		case numero == 5:
			diaDaSemana = "Quinta-feira"
		case numero == 6:
			diaDaSemana = "Sexta-feira"
		case numero == 7:
			diaDaSemana = "Sabado"
		default:
			diaDaSemana = "Número inválido"
		}

		return diaDaSemana
	}


func main() {
	fmt.Println("Switch")

	dia := diaDaSemana(5)
	fmt.Println(dia)

	dia2 := diaDaSemana2(7)
	fmt.Println(dia2)

	dia3 := diaDaSemana3(1)
	fmt.Println(dia3)
}
