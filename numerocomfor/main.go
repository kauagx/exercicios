package main

import "fmt"

func main() {
	var numeroMaior, numeroMenor int
	for {
		fmt.Print("Digite um número maior que 100: ")
		fmt.Scan(&numeroMaior)

		if numeroMaior >= 100 {
			break
		}
		fmt.Println("Peço que digite corretamente o que foi solicitado: ")
	}

	for {
		fmt.Print("Agora digite um número menor que 10: ")
		fmt.Scan(&numeroMenor)

		if numeroMenor < 10 {
			break
		}
		fmt.Print("Peço que digite corretamente o que foi solicitado: ")
	}

	divisao := numeroMaior / numeroMenor
	resto := numeroMaior % numeroMenor

	if resto == 0 {
		fmt.Printf("O número %d cabe %dx no número %d\n", numeroMenor, divisao, numeroMaior)
	} else {
		fmt.Printf("O número %d cabe aproximadamente %dx no número %d\n", numeroMenor, divisao, numeroMaior)
	}
}
