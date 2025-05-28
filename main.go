package main

import "fmt"

func main(){
	var nome string
	var idade int

	fmt.Print("Digite nome e idade:")
	fmt.Scan(&nome, &idade)
	fmt.Println("Nome:", nome)
	fmt.Println("Idade:", idade)
	
	if idade >= 18{
		fmt.Print("Status: Maior de idade")
	} else {
		fmt.Print("Status:Menor de idade")
	}

}