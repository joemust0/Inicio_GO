package main

import "fmt"

func menu() {
	fmt.Println("Qual operação deseja fazer?")
	fmt.Println("1. Soma")
	fmt.Println("2. Subtração")
	fmt.Println("3. Multiplicação")
	fmt.Println("4. Divisão")
}

func main() {
	var num1, num2, resposta int

	fmt.Println("Vamos fazer o cálculo!")

	fmt.Println("Digite o primeiro valor:")
	fmt.Scanln(&num1)

	fmt.Println("Digite o segundo valor:")
	fmt.Scanln(&num2)

	menu()

	fmt.Scanln(&resposta)

	switch resposta {
	case 1:
		fmt.Println("Resultado da soma:", num1+num2)
	case 2:
		fmt.Println("Resultado da subtração:", num1-num2)
	case 3:
		fmt.Println("Resultado da multiplicação:", num1*num2)
	case 4:
		if num2 != 0 {
			fmt.Println("Resultado da divisão:", float64(num1)/float64(num2))
		} else {
			fmt.Println("Não é possível dividir por zero.")
		}
	default:
		fmt.Println("Número inválido")
	}
}
