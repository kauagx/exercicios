package main

import "fmt"

func main() {
	var nota1, nota2 float64

	fmt.Println("Deseja calcular a média de qual matéria?: ")
	var materia string
	fmt.Scan(&materia)

	fmt.Println("Qual foi sua primeira nota?: ")
	fmt.Scan(&nota1)

	fmt.Println("Qual foi sua segunda nota?: ")
	fmt.Scan(&nota2)

	media := (nota1 + nota2) / 2

	var conceito string
	var situacao string

	if media >= 9 {
		conceito = "A"
		situacao = "APROVADO"
	} else if media >= 7 {
		conceito = "B"
		situacao = "APROVADO"
	} else if media >= 5 {
		conceito = "C"
		situacao = "REPROVADO"
	} else if media >= 3 {
		conceito = "D"
		situacao = "REPROVADO"
	} else if media >= 0 {
		conceito = "E"
		situacao = "REPROVADO"
	} else {
		fmt.Println("Média inválida")
	}

	// Exibição dos resultados
	fmt.Printf("\nMatéria: %s\n", materia)
	fmt.Printf("Nota 1: %.2f\n", nota1)
	fmt.Printf("Nota 2: %.2f\n", nota2)
	fmt.Printf("Média: %.2f\n", media)
	fmt.Printf("Conceito: %s\n", conceito)
	fmt.Printf("Status: %s\n", situacao)

}
