package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){

	mapaCursos := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin) // Vai ler os dados que vem do console
	curso := ""

	for curso != "q" {
		var cargaHoraria int
		fmt.Print("Digite um curso ou 'q' para sair: ")
		scanner.Scan()
		curso = scanner.Text() // Retorna toda informação do scan
		if curso != "q" {
			fmt.Print("Digite a carga horario do curso: ")
			fmt.Scanf("%d", &cargaHoraria)

			// Adicionar infos no mapa
			mapaCursos[curso] = cargaHoraria
		}
	}
	fmt.Println("Cursos registrados: ")
	for cursoNome, carga := range mapaCursos {
		fmt.Println("Nome: ", cursoNome, " Carga: ", carga)
	}

	// REMOVENDO ITENS
	curso = ""
	for curso != "q" {
		fmt.Print("Digite o nome do curso a ser excluido ou q para cancelar: ")
		curso = scanner.Text()
		if curso != "q" {
			// Neste get ele retorna o valor do elemento e um booleano dizendo se a chave existe ou não:
			carga, cursoExiste := mapaCursos[curso]
			if cursoExiste {
				delete(mapaCursos, curso)
				fmt.Printf("O curso %s com %dh foi removido", curso, carga)
			} else {
				fmt.Printf("O curso %s não existe", curso)
			}
		}
	}

	fmt.Println("Cursos registrados: ")
	for cursoNome, carga := range mapaCursos {
		fmt.Println("Nome: ", cursoNome, " Carga: ", carga)
	}
}