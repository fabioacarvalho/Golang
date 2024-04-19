package main

import ( 
	"fmt"
	"errors"
)

type Acelerador interface {
	acelerar() error
}

type Freador interface {
	frear() error
}

type Carro struct {

	modelo string
	marca string
	velocidade float32

}

func (carro *Carro) acelerar() error {
	if carro.velocidade < 15 {
		carro.velocidade += 5
		return nil
	} else {
		return errors.New("o carro já atingiu a velocidade máxima")
	}
}

func (carro *Carro) frear() error {
	if carro.velocidade > 0 {
		carro.velocidade -= 5
		return nil
	} else {
		return errors.New("o carro já está parado")
	}
}

func main(){
	
	// Como se fosse o construtor
	carro := Carro { modelo: "Palio", marca: "Fiat"}

	fmt.Println("Iniciando o trajeto")

	opcao := 0

	for opcao != 3 {

		fmt.Print("O que deseja fazer \n 1 - Acelerar \n 2 - Frear \n 3 - Sair \n Escolha uma opção: ")
		fmt.Scanf("%d", &opcao)
		var err error = nil
		switch opcao {
		case 1:
			err = fazerAcelerar(&carro)
		case 2:
			err = fazerFrear(&carro)
		}
		if err != nil {
			fmt.Printf("Não foi possível executar a ação \n")
		} else if opcao != 3 {
			fmt.Printf("O carro está %.2f Km/h \n", carro.velocidade)
		}
	}
}

func fazerAcelerar(veiculo Acelerador) error {
	return veiculo.acelerar()
}

func fazerFrear(veiculo Freador) error {
	return veiculo.frear()
}