package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numeroSecreto := rand.Intn(100) + 1

	const (
		nivel1 = 15
		nivel2 = 10
		nivel3 = 5
	)

	fmt.Println("Escolha o nível de dificuldade do jogo (nivel1, nivel2, nivel3): ")
	var nivelDificuldade string
	fmt.Scan(&nivelDificuldade)

	//Criando a váriavel para armazenar as tentantivas de cada nivel
	var maxTentativas int

	if nivelDificuldade == "nivel1" {
		maxTentativas = nivel1
	} else if nivelDificuldade == "nivel2" {
		maxTentativas = nivel2
	} else if nivelDificuldade == "nivel3" {
		maxTentativas = nivel3
	} else {
		fmt.Println("Nível inválido!")
		return
	}

	fmt.Println("Neste programa você terá que adivinhar o número armazenado")

	//Loop das tentativas até acertar
	for i := 0; i <= maxTentativas; i++ {
		fmt.Println("Tentativa", i, "de", maxTentativas)
		fmt.Print("Digite um número: ")
		var tentativa int
		fmt.Scan(&tentativa)

		if tentativa < numeroSecreto {
			fmt.Println("Você errou, o número é maior")
		} else if tentativa > numeroSecreto {
			fmt.Println("Você errou, o número é menor")
		} else {
			fmt.Println("Parabéns, você acertou o número escolhido!")
			return
		}
	}

	fmt.Println("Você não acertou!. O número secreto era:", numeroSecreto)
}
