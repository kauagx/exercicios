package main

import "fmt"

func main() {
	var idade int
	var nome string

	fmt.Print("Qual o seu nome?: ")
	fmt.Scan(&nome)

	fmt.Printf("Ok, %s.\nAgora me dize qual sua idade: ", nome)
	fmt.Scan(&idade)

	novaIdade := idade + 1

	geral(nome, novaIdade)
}

func geral(nome string, novaIdade int) {
	fmt.Println(nome + ", no próximo aniversário você terá", novaIdade, "anos")
}
