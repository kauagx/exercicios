package main

import "fmt"

func main() {
	var salarioHora, horasSemana float64

	fmt.Println("Qual seu salário por hora?: ")
	fmt.Scan(&salarioHora)

	fmt.Println("Quantas horas você trabalha por semana?: ")
	fmt.Scan(&horasSemana)

	horasTrabalhadas := horasSemana * 5
	salarioBruto := salarioHora * horasTrabalhadas
	impostoRenda := salarioBruto * 11 / 100
	impostoInss := salarioBruto * 8 / 100
	sindicato := salarioBruto * 5 / 100
	descontos := impostoRenda + impostoInss +
		sindicato
	salarioLiquido := salarioBruto - descontos

	fmt.Printf(`Salário Bruto : R$%.2f
IR (11%%) : R$%.2f
INSS (8%%) : R$%.2f
Sindicato (5%%) : R$%.2f
Salário Liquido : R$%.2f `, salarioBruto, impostoRenda, impostoInss, sindicato, salarioLiquido)
}
