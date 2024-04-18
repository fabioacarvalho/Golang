package main

import (
	"fmt"
)

func main(){

	var idade int
	fmt.Printf("Digite a sua idade: ")
	fmt.Scanf("%d", &idade)

	if idade >= 18 {
		fmt.Println("Você é MAIOR de idade.")
	} else {
		fmt.Println("Você é MENOR de idade.")
	}

	var numero int
	fmt.Printf("Digite um número: ")
	fmt.Scanf("%d", &numero)
	if numero < 0 {
		fmt.Println("Este número é negativo")
	} else if numero < 10 {
		fmt.Println("Este número é positivo e contém 1 digito")
	} else {
		fmt.Println("Este número é positivo e contém mais de 1 digito")
	}

}