package main

import "fmt"


var nomeEscola string = "Escola Técnica SENAI"

func main(){

	nome := "Leonardo"
	idade := 16

	fmt.Print(welcome(nome))

	fmt.Println(verificamaioridade(idade))
}