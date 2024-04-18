package main

import (
	"fmt"
)

func main(){

	var numero int
	fmt.Printf("Digite um número: ")
	fmt.Scanf("%d", &numero)

	// FOR

	for i := 0; i <= 10; i++ {
		fmt.Printf("%d x %d = %d \n", numero, i, numero * i)
	}

	// 'WHILE' > Dentro do Go não temos o WHILE mas podemos representar ele utilizando o FOR da seguinte forma:
	fmt.Printf("\n TABUADA COM WHILE \n")
	i := 0
	for i <= 10 {
		fmt.Printf("%d x %d = %d \n", numero, i, numero * i)
		i++
	}

}