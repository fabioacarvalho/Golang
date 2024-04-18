package main

import (
	"fmt"
	"container/list"
)

func main(){

	numeros := list.New() // Inicializando uma lista
	el := numeros.PushBack(1) // Inserir um elemento na parte de trás da lista

	numeros.PushFront(0)
	numeros.PushBack(2)

	for e := numeros.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	fmt.Println("---------")

	numeros.Remove(el)

	for e := numeros.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}


