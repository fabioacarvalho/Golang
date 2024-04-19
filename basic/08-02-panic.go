package main

import ( 
	"fmt"
)

func main(){
	fmt.Println("Bem-Vindo!")
	defer fmt.Println("Obrigado por utilizar nosso sistema!")

	fmt.Print("Digite um número maior que 10: ")
	var numero int
	fmt.Scanf("%d", &numero)

	if numero < 10 {
		panic("Número invalido")
	} else {
		fmt.Println("Bom numero...")
	}

}
