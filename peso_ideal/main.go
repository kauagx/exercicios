package main

import "fmt"

func main() {
	var altura, peso, IMC float32
	var sexo string

	fmt.Print("Qual seu sexo? [M/F]: ")
	fmt.Scan(&sexo)

	/*
	O || retorna verdadeiro se pelo menos uma condição for verdadeira
	*/
	if sexo == "M" || sexo == "F" {
	} else {
		fmt.Println("Entrada inválida! Reinicie o programa e digite 'M' ou 'F'.")
		return
	}

	fmt.Print("Qual a sua altura (m)?: ")
	fmt.Scan(&altura)

	fmt.Print("Qual o seu peso (kg)?: ")
	fmt.Scan(&peso)

	// Cálculo do IMC
	IMC = peso / (altura * altura)

	fmt.Printf("Seu IMC é: %.2f\n", IMC)

	
	if sexo == "M" {
		if IMC < 20 {
			fmt.Println("Você está abaixo do peso.")
		} else if IMC >= 20 && IMC < 25 {
			fmt.Println("Seu peso está normal.")
		} else if IMC >= 25 && IMC < 30 {
			fmt.Println("Você está com sobrepeso.")
		} else {
			fmt.Println("Você está com obesidade.")
		}
	} else if sexo == "F" { 
		if IMC < 19 {
			fmt.Println("Você está abaixo do peso.")
		} else if IMC >= 19 && IMC < 24 {
			fmt.Println("Seu peso está normal.")
		} else if IMC >= 24 && IMC < 29 {
			fmt.Println("Você está com sobrepeso.")
		} else {
			fmt.Println("Você está com obesidade.")
		}
	}
}