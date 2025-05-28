package main

import "fmt"

func main(){
	var a, b int

	fmt.Print("Digite dois números")
	fmt.Scan(&a, &b)
	fmt.Println("Os números são:", a, b)
	fmt.Println("Soma:", a+b)
	fmt.Println("Subtração:", a-b)
	fmt.Println("Multiplicação:", a*b)
	fmt.Println("Divisão:", a/b)
	fmt.Println("Resto:", a%b)

	a += 1
	fmt.Println("Incremento de a:", a)
}