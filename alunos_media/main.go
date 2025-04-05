package main

import "fmt"

type Aluno struct {
	nome string
	p1   float64
	p2   float64
}

func main() {
	alunos := []Aluno{}
	fmt.Print("Quantos alunos tem na sala?: ")
	var quantidade int
	fmt.Scan(&quantidade)

	for i := range quantidade {
		fmt.Printf("Qual o nome do %d aluno: ", i+1)
		var nome string
		fmt.Scan(&nome)

		fmt.Print("Qual a primeira nota do aluno?: ")
		var nota1 float64
		fmt.Scan(&nota1)

		fmt.Print("Qual a segunda nota do aluno?: ")
		var nota2 float64
		fmt.Scan(&nota2)

		alunos = append(alunos, Aluno{nome, nota1, nota2})
	}
	for _, aluno := range alunos {
		media := (aluno.p1 + aluno.p2) / 2
		fmt.Printf("%s p1:%.2f p2:%.2f media:%.2f\n", aluno.nome, aluno.p1, aluno.p2, media)
	}

}
