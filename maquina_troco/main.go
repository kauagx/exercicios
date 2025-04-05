package main

import "fmt"

func main() {
	// Lista de cédulas e moedas disponíveis
	valores := []float64{100, 50, 20, 10, 5, 2, 1, 0.50, 0.25, 0.10, 0.05, 0.01}

	fmt.Println("Qual o valor do produto a ser comprado?: ")
	var produto float64
	fmt.Scan(&produto)

	fmt.Println("Qual o valor a ser pago?: ")
	var pagamento float64
	fmt.Scan(&pagamento)

	// Verifica se o pagamento é suficiente
	if pagamento < produto {
		fmt.Println("Pagamento insuficiente!")
		return
	}

	// Calcula o troco corretamente
	troco := pagamento - produto

	fmt.Printf("O valor pago foi de R$%.2f\n", pagamento)
	fmt.Printf("O preço do produto é R$%.2f\n", produto)
	fmt.Printf("O troco é de R$%.2f\n", troco)

	// Distribui as notas e moedas
	for _, valor := range valores {
		quantidade := int(troco / valor)
		if quantidade > 0 {
			if valor >= 1 {
				fmt.Printf("%d nota(s) de R$%d\n", quantidade, int(valor))
			} else {
				fmt.Printf("%d moeda(s) de R$%.2f\n", quantidade, valor)
			}
			troco -= float64(quantidade) * valor
		}
	}
}
