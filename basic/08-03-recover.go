package main

import ( 
	"fmt"
)

func main(){
	
	fmt.Println("Bem-Vindo")

	// Função anônima
	defer func(){
		if ex := recover(); ex != nil {
			fmt.Printf("Desculpe, ocorreu um erro %s", ex)
		}
		fmt.Println("Obrigado por utilizar nosso software")
	}()

	fmt.Print("Digite um número maior que 10: ")
	var numero int
	fmt.Scanf("%d", &numero)

	if numero < 10 {
		panic("Número invalido")
	} else {
		fmt.Println("Bom numero...")
	}
}