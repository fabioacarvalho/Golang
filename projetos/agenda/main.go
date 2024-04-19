package main

import ( 
	"fmt"
	"os"
	"bufio"
	"strings"
	"time"
)

// ARQUIVO : indica o local onde os contatos serão salvos
const ARQUIVO string = "agenda.txt"

// ConversivelParaString : interface que especifica como um item pode ser convertido para uma string
type ConversivelParaString interface {
	toString() string
}

// Contato : estrutura que representa um contato
type Contato struct {
	nome string
	formaContato string
	valorContato string
}

func (contato *Contato) toString() string {
	return fmt.Sprintf("%s|%s|%s \n", contato.nome, contato.formaContato, contato.valorContato)
}

// GerenciadorContatos : componente responsavel por gerenciar os contatos
type GerenciadorContatos struct {}

func (gerenciador *GerenciadorContatos) carregarContatos() ([]Contato, error) {
	contatos := make([]Contato, 0)
	if _, err := os.Stat(ARQUIVO); !os.IsNotExist(err) {
		arquivo, erro := os.Open(ARQUIVO)
		if erro != nil {
			return contatos, err
		}
		defer arquivo.Close()

		scanner := bufio.NewScanner(arquivo)
		for scanner.Scan() {
			linhaContato := scanner.Text()
			infoContato := strings.Split(linhaContato, "|")

			contatos = append(contatos, Contato{nome: infoContato[0], formaContato: infoContato[1], valorContato: infoContato[2]})
		}
	}
	return contatos, nil
}

func (gerenciador *GerenciadorContatos) salvarContato(contato ConversivelParaString) error {
	arquivo, err := os.OpenFile(ARQUIVO, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer arquivo.Close()

	if _, e := arquivo.WriteString(contato.toString()); e != nil {
		return e
	}
	return nil
}


func main(){
	gerenciador := GerenciadorContatos{}
	opcao := 0

	for true {
		fmt.Print("O que deseja fazer: \n 1-Listar contatos \n 2-Criar contato \n 3-Sair \n")
		fmt.Scanf("%d", &opcao)
		switch opcao {
		case 1:
			listarContatos(&gerenciador)
		case 2:
			criarNovoContato(&gerenciador)
		}
		if opcao == 3 {
			break
		}
	}
	fmt.Println("Bye!")
}

func listarContatos(gerenciador *GerenciadorContatos) {
	contatos, err := gerenciador.carregarContatos()
	if err != nil {
		fmt.Printf("Não foi pósssível carregar os contatos: %s \n", err)
	} else {
		fmt.Println("Lista de contatos: \n")
		for _, contato := range contatos {
			fmt.Printf("  - %s, %s: %s", contato.nome, contato.formaContato, contato.valorContato)
		}
	}
}

func criarNovoContato(gerenciador *GerenciadorContatos){
	novoContato := Contato{}
	fmt.Print("Nome do contato: \n")
	fmt.Scanf("%s", &novoContato.nome)
	time.Sleep(2)
	fmt.Print("Tipo de contato: \n")
	fmt.Scanf("%s", &novoContato.formaContato)
	time.Sleep(2)
	fmt.Print("Valor do contato: \n")
	fmt.Scanf("%s", &novoContato.valorContato)
	time.Sleep(2)

	err := gerenciador.salvarContato(&novoContato)
	if err != nil {
		fmt.Printf("Houve um erro ao salvar o contato %s \n", err)
	}
}