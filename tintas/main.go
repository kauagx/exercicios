package main

import (
	"fmt"
	"math"
)

func main() {
	var areaPintada float64
	capacidadeLata := 18.0
	cobertura := 3.0
	valorLata := 80.0

	fmt.Print("Qual o tamanho da área que vai ser pintada em metros quadrados?: ")
	fmt.Scan(&areaPintada)

	litros := areaPintada / cobertura
	latas := math.Ceil(litros / capacidadeLata)
	precoTotal := latas * valorLata

	fmt.Printf("Quantidade de latas necessárias: %.0f\n", latas)
	fmt.Printf("Valor a ser gasto: R$%.2f\n", precoTotal)
}
