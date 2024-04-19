package main
import (
	"fmt"
)

func main(){

	numero1 := 5
	numero2 := 2
	operacao := "/"

	var metodoOperacao func(n1 int, n2 int) (int, int)
	if operacao == "+" {
		metodoOperacao = func(n1, n2 int) (int, int) {
			return n1 + n2, 0
		}
	} else if operacao == "-" {
		metodoOperacao = func(n1, n2 int) (int, int) {
			return n1 - n2, 0
		}
	} else if operacao == "/" {
		metodoOperacao = func(n1, n2 int) (int, int) {
			return n1 / n2, n1 % n2
		}
	}
	resultado1, resultado2 := metodoOperacao(numero1, numero2)
	fmt.Printf("%d %s %d = %d, %d", numero1, operacao, numero2, resultado1, resultado2)
}
