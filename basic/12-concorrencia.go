package main

import ( 
	"fmt"
)

func main(){
	var limite int

	canal1 := make(chan int)
	canal2 := make(chan int)

	fmt.Print("Informe um limite: ")
	fmt.Scanf("%d", &limite)

	// Go routine torna ela assincrona, ou seja, executa em uma thread diferente
	go conteAte(limite, canal1)

	go func(n int){
		resultado := 0
		for i :=0; i <= n*10; i++ {
			fmt.Printf(" - [anonimo] O número é %d \n", i)
			resultado = i * 10
		}

		canal2 <- resultado
	}(limite)

	for i :=0; i <= limite*10; i++ {
		fmt.Printf(" - [main] O número é %d \n", i)
	}

	// Recebendo os valores dos canais
	resultadoCanal1 := <- canal1
	resultadoCanal2 := <- canal2
	fmt.Printf(" Canal 1: %d \n Canal 2: %d", resultadoCanal1, resultadoCanal2)
}

func conteAte(limite int, canal chan int){
	resultado := 0
	for i :=0; i <= limite*20; i++ {
		fmt.Printf(" - [conteAte] O número é %d \n", i)
		resultado = i * 20
	}
	canal <- resultado
}