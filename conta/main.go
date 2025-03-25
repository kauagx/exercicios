package main

import "fmt"

func main() {
	var valorConta float32
	var totalPessoas int

	fmt.Print("Qual o valor total da conta?: ")
	fmt.Scan(&valorConta)

	fmt.Print("Certo, ficou no total de R$", valorConta,".\nAgora me diz qual o total de pessoas?: ")
	fmt.Scan(&totalPessoas)

	totalCada := valorConta / float32(totalPessoas) 

	valorCada(totalCada)
}	
func valorCada(totalCada float32) {
	fmt.Printf("Ficou R$%.2f pra cada um\n", totalCada)
}