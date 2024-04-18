package main

import (
	"fmt"
)


func main(){
	fmt.Println("INTEIROS SEM SINAL")
	var u1 byte = 255
	var u2 uint16 = 5555
	var u3 uint32 = 5555
	var u4 uint64 = 5555
	fmt.Println(u1)
	fmt.Println(u2)
	fmt.Println(u3)
	fmt.Println(u4)
	
	fmt.Println("INTEIROS COM SINAL")
	var i1 int8 = 127
	var i2 int16 = 1000
	fmt.Println(i1)
	fmt.Println(i2)
	
	var t1 uint = 1010
	fmt.Println(t1)
	
	var t2 int = 2000
	fmt.Println(t2)

	fmt.Println("PONTOS FLUTUANTES")
	var f1 float32 = 1
	fmt.Println(f1)
	var f2 float64 = 2
	fmt.Println(f2)
	var f3 complex64 = 3
	fmt.Println(f3)

	fmt.Println("STRINGS")
	var s1 string = "String 1"
	fmt.Println(s1)

	var s2 string = "String \n 2"
	fmt.Println(s2)

	var s3 string = `String
Quebra de Linha
	`
	fmt.Println(s3)

	fmt.Println("Valor Unicode: ")
	fmt.Println(s1[0]) // Valor unicode

	fmt.Println("BOOLEANOS")
	var b1 bool = true
	fmt.Println(b1)

	fmt.Println("MULTIPLAS DECLARACOES")
	var nome, sobrenome string = "Fabio", "Carvalho"
	fmt.Println(nome + " " + sobrenome)
	
	var (
		nome2 string = "Fabio"
		sobrenome2 string = "Carvalho"
		idade int = 29
	)
	
	fmt.Println(nome2 + " " + sobrenome2 + " ", idade, " anos.")

	fmt.Println("INFERENCIA DE TIPOS")
	var nome = "Jose Campos"
	fmt.Println(nome)
	
	nomeCompleto := "Jose dos Campos Oliveira"
	fmt.Println(nomeCompleto)

	fmt.Println("CONSTANTES")
	const slogan string = "Seja bem-vindo"
	const slogan2 = "Seja bem-vindo2"
	const (
		slogan3 = "Seja bem-vindo3"
		valor = 39.9
		qtde = 5
	)
	fmt.Println(slogan)
	fmt.Println(slogan2)
	fmt.Println(slogan3 + " O valor total ficou em: ", qtde * valor)

}

