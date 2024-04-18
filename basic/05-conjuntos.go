package main

import (
	"fmt"
)

func main(){

	var amigos [5]string
	fmt.Println(amigos)

	for i := 0; i < len(amigos); i++ {
		fmt.Print("Digine o nome de um dos seus amigos: ")
		fmt.Scanf("%s", &amigos[i])
	}
	fmt.Println(amigos)

	// Um tipo de "Foreach" > aqui é passado o indice e o elemento.
	for _, amigo := range amigos {
		fmt.Printf("- %s \n", amigo)
	}

	arrayInicializado := [3]int {1, 2, 3}
	fmt.Println(arrayInicializado)

	var matriz [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println("Digite um número: ")
			fmt.Scanf("%d", &matriz[i][j])
		}
	}
	fmt.Println(matriz)
}