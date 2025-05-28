package main


func welcome(nome string) string{
	return "Olá " + nome + " Você está na" + nomeEscola
}

func verificamaioridade(idade int) string{
	if idade >= 18 {
		return "Você é maior de idade"
	} else {
		return "Você é menor de idade"
	}
}