package main
import ( 
	"fmt"
)
func main(){
	fmt.Println("Estou abrindo a conexão com o banco de dados.")

	// Ele sempre será executado no final. Vale lembrar que o DEFER funciona no modelo LIFO.
	defer fecharConsulta() // Segundo a ser executado
	defer fmt.Println("Vou executar a função") // Primeiro a ser executado
	executarConsulta()
}

func executarConsulta(){
	fmt.Println("Estou executando a conexão com o banco de dados.")
}

func fecharConsulta(){
	fmt.Println("Estou fechando a conexão com o banco de dados.")
}