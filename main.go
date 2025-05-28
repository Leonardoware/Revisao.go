package main

import "fmt"

func main(){
	Produtos := make(map[string]int)

	Produtos["Coxinha"] += 10
	Produtos["Pão de Queijo"] += 15
	Produtos["Refrigerante"] += 20

	for produto, quantidade := range Produtos{
		fmt.Println(produto, quantidade)
	}

	fmt.Println("\nApós a venda:\n")
	
	posVenda := make(map[string]int)
	
	posVenda["Coxinha"] += 8
	posVenda["Pão de Quejo"] += 14
	posVenda["Refrigerante"]+= 20

	for produtos2, quantiaAtual := range posVenda{
		fmt.Println(produtos2, quantiaAtual)
	}
}