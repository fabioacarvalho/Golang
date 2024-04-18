package main

import (
	"fmt"
)

func main(){

	// Passamos o [] e o que iremos armazenar e seu tamanho inicial
	amigos := make([]string, 3)
	nome := ""
	i := 0

	for nome != "q" {
		fmt.Print("Digite o nome de um amigo ou 'q' para parar: ")
		fmt.Scanf("%d", &nome)
		if nome != "q" {
			if i < 3 {
				amigos[i] = nome
			} else {
				amigos = append(amigos, nome)
			}
			i++
		}
	}
	fmt.Println(amigos)
	fmt.Printf("O tamanho atual do slice é de: ", len(amigos))

	outroSlice := amigos[1:3]
	fmt.Println(outroSlice)
}