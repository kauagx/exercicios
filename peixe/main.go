package main

import "fmt"

func main() {
	var peso float64
	const (
		limite        = 50
		multaAplicada = 4
	)

	fmt.Println("João, quantos kilos de peixe vc pescou?: ")
	fmt.Scan(&peso)

	if peso > limite {
		excesso := peso - limite
		multa := excesso * multaAplicada
		fmt.Printf("A multa aplicada é de R$%.2f", multa)
	} else {
		fmt.Print("Hoje não deverá pagar multa")
	}
}
