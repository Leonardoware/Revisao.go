package main

import "fmt"

func main(){
	var usuario string
	var senha string

	usuarioCorreto := "Leonardo"
	senhaCorreta := "3005"

	fmt.Print("Digite seu usuário:")
	fmt.Scan(&usuario)
	fmt.Print("Digite sua senha:")
	fmt.Scan(&senha)

	if usuario == usuarioCorreto && senha == senhaCorreta{
		fmt.Println("Conta acessada com sucesso!")
	} else{
		fmt.Println("Seu nome de usuário ou senha estão incorretos, por favor, tente novamente")
	}

	
}