package main

import (
	"fmt"
	"math/rand/v2"
	"os"
)

const (
	letrasMinusculas = "abcdefghijklmnopqrstuvwxz"
	letrasMaiusculas = "ABCDEFGHIJKLMNOPQRSTUVWXZ"
	numeros          = "0123456789"
)

func main() {

	fmt.Println("Informe a quantidade de caracteres da senha:")
	var quantidade int
	fmt.Scan(&quantidade)

	fmt.Println("Deve incluir minúsculas s/n?")
	var incluirMinusculas string
	fmt.Scan(&incluirMinusculas)

	fmt.Println("Deve incluir maiúsculas s/n?")
	var incluirMaiusculas string
	fmt.Scan(&incluirMaiusculas)

	fmt.Println("Deve incluir números s/n?")
	var incluirNumeros string
	fmt.Scan(&incluirNumeros)

	caracteres := ""
	if incluirMinusculas == "s" {
		caracteres += letrasMinusculas
	}
	if incluirMaiusculas == "s" {
		caracteres += letrasMaiusculas
	}
	if incluirNumeros == "s" {
		caracteres += numeros
	}

	if caracteres == "" {
		fmt.Println("Você precisa informar pelo menos 1 tipo de caracter")
		os.Exit(1)
	}

	senha := ""

	for i := 0; i < quantidade; i++ {
		numero := rand.IntN(len(caracteres))
		//fmt.Println(numero)
		letra := string(caracteres[numero])
		//fmt.Println(letra)
		senha = senha + letra
		//fmt.Println(senha)
	}

	fmt.Println(senha)
}
