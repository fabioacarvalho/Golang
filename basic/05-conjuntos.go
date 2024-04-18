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
}