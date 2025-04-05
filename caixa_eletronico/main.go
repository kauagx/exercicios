package main

import "fmt"

func main() {

	cedulas := []float64{100.0, 50.0, 20.0, 10.0, 5.0, 1.0} // Notas organizadas do maior para o menor
	min, max := 10.0, 600.0

	fmt.Println("\nQual o valor que será sacado?: ")
	var saque float64
	fmt.Scan(&saque)

	// Validação do valor do saque
	for saque > max || saque < min {
		fmt.Println("Valor inválido! O saque deve estar entre R$10 e R$600")
		fmt.Println("Digite um novo valor para saque: ")
		fmt.Scan(&saque)
	}

	fmt.Println("\nOBRIGADO PELA INFORMAÇÃO. AGORA VAMOS SEGUIR COM A OPERAÇÃO")

	// Loop para calcular as cédulas necessárias
	fmt.Println("Notas fornecidas:")
	for _, cedula := range cedulas {
		quantidade := int(saque) / int(cedula)
		if quantidade > 0 {
			fmt.Printf("%d nota(s) de R$%.2f\n", quantidade, cedula)
			saque = float64(int(saque) % int(cedula)) // explicação float64(int(saque) "o float64 converte o valor dentro do parentes"
		}
	}
}
