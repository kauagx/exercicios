package main

import "fmt"

func main() {
	var numeroMaior, numeroMenor int
	fmt.Print("Digite um número maior que 100: ")
	fmt.Scan(&numeroMaior)

	if numeroMaior < 100 {
		fmt.Println("O número deve ser maior que 100!")
		return
	} 

	fmt.Print("Agora digite um número menor que 10: ")
	fmt.Scan(&numeroMenor)

	if numeroMenor > 10 {
		fmt.Println("O número tem que ser menor que 10!")
		return
	}

	divisao := numeroMaior / numeroMenor
	resto := numeroMaior % numeroMenor

	if resto == 0 {
		fmt.Printf("O número %d cabe %dx no número %d\n", numeroMenor, divisao, numeroMaior)
	} else {
		fmt.Printf("O número %d cabe aproximadamente %dx no número %d\n", numeroMenor, divisao, numeroMaior)
	}
}

