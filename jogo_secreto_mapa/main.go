package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numeroSecreto := rand.Intn(100) + 1

	dificuldades := map[int]int{
		1: 15,
		2: 10,
		3: 5,
	}

	fmt.Println("Escolha o nível de dificuldade do jogo (1, 2, 3): ")
	var nivelDificuldade int
	fmt.Scan(&nivelDificuldade)

	//Criando a váriavel para armazenar as tentantivas de cada nivel
	var maxTentativas int

	maxTentativas, ok := dificuldades[nivelDificuldade]
	if !ok {
		fmt.Println("Nível inválido!")
		return
	}

	fmt.Println("Neste programa você terá que adivinhar o número armazenado")

	//Loop das tentativas até acertar
	for i := 1; i <= maxTentativas; i++ {
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

	fmt.Println("Você não acertou! O número secreto era:", numeroSecreto)
}
