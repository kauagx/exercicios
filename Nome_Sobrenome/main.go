package main

import "fmt"

func main() {
    var nome string
    fmt.Print("Digite seu primeiro nome: ")
    fmt.Scan(&nome)

    var sobrenome string
    fmt.Print("Agora digite seu sobrenome: ")
    fmt.Scan(&sobrenome)
    imprimeNome(nome, sobrenome)
}

func imprimeNome(nome string, sobrenome string) {
    fmt.Println("Obrigado pela informação!\nSeu nome é", nome, "e o sobrenome", sobrenome)
}