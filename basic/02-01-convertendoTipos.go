package main

import (
	"fmt"
	"strconv"
)

func main(){

	var texto string

	fmt.Print("Digite um número: ")
	fmt.Scanf("%s", &texto)
	var numero int

	// O _ neste caso seria para ignorar o erro caso o mesmo seja gerado.
	numero, _ = strconv.Atoi(texto) // Essa função retorna o valor ou um erro > (int, error) 
	fmt.Println(numero)

	// Outra forma de converter > strconv.ParseInt(valor_a_ser_convertido, base_de_conversao, unidade_de_conversao)
	numero2, _ := strconv.ParseInt(texto, 10, 32)
	fmt.Println(numero2)
}