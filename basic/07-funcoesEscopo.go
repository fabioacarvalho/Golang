package main

import (
	"fmt"
)

func main(){

	var numero1, numero2 int
	var operacao string
	fmt.Print("Digite o primeiro numéro: ")
	fmt.Scanf("%d", &numero1)
	fmt.Print("Digite a operação (+, -, *, /, $): ")
	fmt.Scanf("%d", &operacao)
	fmt.Print("Digite o segundo numéro: ")
	fmt.Scanf("%d", &numero2)


	if operacao == "+" {
		somar(numero1, numero2)
	} else if operacao == "-" {
		resultado := subtrair(numero1, numero2)
		fmt.Printf("%d + %d = %d", numero1, numero2, resultado)
	} else if operacao == "*" {
		resultado := multiplicar(numero1, numero2)
		fmt.Printf("%d x %d = %d", numero1, numero2, resultado)
	} else if operacao == "/" {
		resultado, resto := dividir(numero1, numero2)
		fmt.Printf("%d / %d = %d", numero1, numero2, resultado)
		fmt.Printf("%d %% %d = %d", numero1, numero2, resto)
	}

}

func somar(n1 int, n2 int) {
	fmt.Printf("%d + %d = %d", n1, n2, n1 + n2)
}

func subtrair(n1 int, n2 int) int {
	return n1 - n2
}

// Aqui já podemos definir a variavel que será retornada
func multiplicar(n1 int, n2 int) (resultado int) {
	resultado = n1 * n2
	return
}

// Retornando multiplos valores
func dividir(n1 int, n2 int)(int, int){
	quociente := n1 / n2
	resto := n1 % n2
	return quociente, resto
}
