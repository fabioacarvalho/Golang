package main

import ( 
	"fmt"
)

func main(){
	
	var saldo2 *float64 = new(float64)
	var saldo float64
	fmt.Print("Digite seu saldo: ")
	fmt.Scanf("%f", &saldo)
	fmt.Print("Digite seu saldo 2: ")
	fmt.Scanf("%f", saldo2) // Aqui eu não preciso do & pois o saldo2 já é a posição na memória

	calcularRendimento(&saldo)
	fmt.Printf("Seu saldo com rendimento é %.2f \n", saldo)
}

func calcularRendimento(saldo *float64) {
	*saldo += *saldo * 0.03
}