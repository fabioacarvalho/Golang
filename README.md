# Golang
 
## Summary

- Basic
    - 1 - Preparação do Ambiente e Conceitos Iniciais
    - 2 - Variáveis e Tipos de Dados
    - 3 - Operadores Aritméticos
    - 4 - Estruturas Condicionais
    - 5 - Estruturas de repetição
	- 6 - Conjuntos
	- 7 - Funções e Escopo

# Basic

## 1 - Preparação do Ambiente e Conceitos Iniciais

Basta acessar o site e baixar o pacote de instalação. a IDE mais utilizada é o VSCODE apesar da JetBrains estar desenvolvendo uma própria para o Go mas ainda está em fase beta nesta data.

Para fazer um import no GO eu posso utilizar da seguinte forma:

```go
	package main
	import (
		"fmt"
	)
```

Para isso é muito simples, basta utilizar a palavra `func` seguindo por `(){}` ficando:

```go
	package main

	import (
	    "fmt"
	)
	
	func main(){
	    fmt.Println("Ola mundo!")
	}
```

Para executar o código basta passar no terminal o seguinte comando:
```shell
go run nome_do_arquivo.go
```


## 2 - Variáveis e Tipos de Dados

Dentro do Go nós temos o que chamamos de **short initialize** ou [[Inferência de Tipos]].
## Inteiros

```go
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
}
```

## Floats

```go
package main

import (
	"fmt"
)

func main(){
	fmt.Println("PONTOS FLUTUANTES")
	var f1 float32 = 1
	fmt.Println(f1)
	var f2 float64 = 2
	fmt.Println(f2)
	var f3 complex64 = 3
	fmt.Println(f3)
}
```

## Strings

```go
package main

import (
	"fmt"
)

func main(){
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

}
```


## Booleanos

```go
package main

import (
	"fmt"
)

func main(){
    fmt.Println("BOOLEANOS")
    var b1 bool = true
    fmt.Println(b1)
}
```

## Múltiplas Variáveis

```go
package main

import (
	"fmt"
)

func main(){
    fmt.Println("MULTIPLAS DECLARACOES")
    var nome, sobrenome string = "Fabio", "Carvalho"
    fmt.Println(nome + " " + sobrenome)
    var (
        nome2 string = "Fabio"
        sobrenome2 string = "Carvalho"
        idade int = 29
    )
    fmt.Println(nome2 + " " + sobrenome2 + " ", idade, " anos.")
}
```

Neste item como podemos ver podemos definir múltiplas variáveis apenas abrindo os parênteses. O bom deste método é que podemos passar outros tipos de vars como int ou bool.

## Constantes

Todas as demais aplicações mencionadas acima se aplicam para a const.

```go
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

```



## 3 - Operadores Aritméticos

Operadores comuns.

```go
package main

import (
    "fmt"
)

func main(){
    var numero1 int
    var numero2 int

    fmt.Print("Digite o primeiro número: ")
    fmt.Scanln(&numero1)
    fmt.Print("Digite o segundo número: ")
    fmt.Scanln(&numero2)

    fmt.Printf("%d + %d = %d \n", numero1, numero2, numero1 + numero2)
    fmt.Printf("%d - %d = %d \n", numero1, numero2, numero1 - numero2)
    fmt.Printf("%d * %d = %d \n", numero1, numero2, numero1 * numero2)
    fmt.Printf("%d / %d = %d \n", numero1, numero2, numero1 / numero2)
    fmt.Printf("%d %% %d = %d \n", numero1, numero2, numero1 % numero2)

    incremento := numero1
    incremento += numero2
    fmt.Printf("O incremento de %d com %d é %d", numero1, numero2, incremento)

    fmt.Printf("\n")

    decremento := numero1
    decremento -= numero2
    fmt.Printf("O decremento de %d com %d é %d", numero1, numero2, decremento)
}
```

>Aqui nós passamos o nome da variável e & comercial que indica a posição na memoria do item.

## 4 - Estruturas Condicionais

No Go dentro do `if` não utilizamos os parênteses, a menos que você deseje dar ordem de precedência para execução. Logo a estrutura fica `if valor1 >= valor2 {}.

```go
package main

import (
    "fmt"
)

func main(){

	// Condições normais
    var idade int
    fmt.Printf("Digite a sua idade: ")
    fmt.Scanf("%d", &idade)
    if idade >= 18 {
        fmt.Println("Você é MAIOR de idade.")
    } else {
        fmt.Println("Você é MENOR de idade.")
    }

	// Condições aninhadas

    var numero int
    fmt.Printf("Digite um número: ")
    fmt.Scanf("%d", &numero)
    if numero < 0 {
        fmt.Println("Este número é negativo")
    } else if numero < 10 {
        fmt.Println("Este número é positivo e contém 1 digito")
    } else {
        fmt.Println("Este número é positivo e contém mais de 1 digito")
    }
}
```


### Switch

Dentro do Go assim como outras linguagens temos o `switch` quando possuímos muitos casos a serem analisados.

```go
package main

import (
    "fmt"
)

func main(){
    var mes int
    fmt.Printf("Digite o número do mês: ")
    fmt.Scanf("%d", &mes)

    switch mes {
    case 1:
        fmt.Println("Esse mês é Janeiro")
    case 2:
        fmt.Println("Esse mês é Fevereiro")
    case 3:
        fmt.Println("Esse mês é Março")
    case 4:
        fmt.Println("Esse mês é Abril")
    case 5:
        fmt.Println("Esse mês é Maio")
    case 6:
        fmt.Println("Esse mês é Junho")
    case 7:
        fmt.Println("Esse mês é Julho")
    case 8:
        fmt.Println("Esse mês é Agosto")
    case 9:
        fmt.Println("Esse mês é Setembro")
    case 10:
        fmt.Println("Esse mês é Outubro")
    case 11:
        fmt.Println("Esse mês é Novembro")
    case 12:
        fmt.Println("Esse mês é Dezembro")
    default:
        fmt.Println("Esse mês é inválido")
    }

    switch mes {
    case 1, 2, 3:
        fmt.Println("Primeiro trimestre")
    case 4, 5, 6:
        fmt.Println("Segundo trimestre")
    case 7, 8, 9:
        fmt.Println("Terceiro trimestre")
    case 10, 11, 12:
        fmt.Println("Quarto trimestre")
    }
}
```



## 5 - Estruturas de repetição

### FOR

A maneira tradicional de fazermos o `for` é muito simples, até parecido com a sintax do javascript, no entanto sem utilizar os parênteses.

```go
package main

import (
    "fmt"
)

func main(){
	var numero int
    fmt.Printf("Digite um número: ")
    fmt.Scanf("%d", &numero)

    for i := 0; i <= 10; i++ {
        fmt.Printf("%d x %d = %d \n", numero, i, numero * i)
    }
}
```

> Como precisamos iniciar uma variável dentro do for, aqui podemos utilizar a sintax reduzida de agregação `:=` lembrando que sem a utilização dos `()`.

## Um tipo de 'FOR EACH'

No Go nós não temos o foreach como em outras linguagens, mas podemos utilizar da seguinte forma:

```go
package main

import (
    "fmt"
)

func main(){
    numeros := [3]int {1, 2, 3}
    fmt.Println(amigos)
    // Um tipo de "Foreach" > aqui é passado o indice e o elemento.
    for _, numero := range numeros {
        fmt.Printf("- %s \n", numero)
    }
}
```

### WHILE

Dentro do Go não possuímos o `while` mas podemos reproduzi-lo com o `for` da seguinte forma:

```go
package main

import (
    "fmt"
)

func main(){
    var numero int
    fmt.Printf("Digite um número: ")
    fmt.Scanf("%d", &numero)

    i := 0
    for i <= 10 {
        fmt.Printf("%d x %d = %d \n", numero, i, numero * i)
        i++
    }
}
```


Perceba que a estrutura é bem parecida.


## 6 - Conjuntos

### Vetores

Para inicializar um array a estrutura básica é `var amigos [5]string`, podemos ver que a definição é feita na seguinte ordem:
1. Nome
2. Quantidade
3. Tipo

```go
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

    arrayInicializado := [3]int {1, 2, 3}
    fmt.Println(arrayInicializado)

    var matriz [3][3]int
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            fmt.Println("Digite um número: ")
            fmt.Scanf("%d", &matriz[i][j])
        }
    }
    fmt.Println(matriz)
}
```

No entanto em Go não é tão comum a utilização de vetores pois eles acabam sendo meio inflexíveis e limitados. Por outro lado o Go oferece algo chamado de Slices que veremos a seguir.

### Slices

Para inicializar um slice utilizamos uma função **built-in** chamada `make()` conforme exemplo:

```go
amigos := make([]string, 3)
```

Aqui podemos ver que passamos o tipo de dado que será armazenado no slice e o tamanho inicial. Isso não significa que o tamanho dele será limitado a 3.

Abaixo temos um exemplo de como podemos aumentar o tamanho do nosso slice, digamos que definimos inicialmente 3 posições para o slice e apareça a necessidade de adicionar mais, para isso basta utilizar outra função **built-in** chamada `append()` para adicionar um novo elemento.

A função `append()` recebe os seguintes parâmetros **nome_slice, item_a_ser_adicionado**  por exemplo `amigos = append(amigos, nome)` além disso é possível passar mais de um item para adicionar `amigos = append(amigos, nome, "Jose", "Maria", "Joao")` dessa forma estaria adicionar estes outros elementos, abaixo um exemplo completo:

```go
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

    outroSlice := amigos[1:3] // Inicializando um novo slice apenas com os itens 1 e 2 do slice amigos
    fmt.Println(outroSlice)
}
```

### Dicionários / Mapa

Para criar um dicionário ou um mapa utilizamos a função **built-in** `make(map[])` e definimos o modelo de chave e valor para eles ficando na seguinte estrutura `mapaCursos := make(map[string]int)` sendo que o valor dentro de `[]` é relacionado a chave, e o `[]int` está relacionado ao valor do conteúdo.

```go
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
```


### Listas Duplamente ligadas

Com o método container/list podemos adicionar os valores por último ou por primeiro através dos métodos `Pushback()` e `PushFront()`.
Dentro deste modelo ao interar pela lista podemos ver que o elemento (e) conhece o seu anterior e proximo elemento, isso se chama de lista duplamente ligada. 


```go
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
```

## 7 - Funções e Escopo

Para criarmos um função ou método em Go devemos sempre iniciar com a keyword **func** logo a estrutura padrão será:

```go
func somar(n1 int, n2 int){
    fmt.Printf("%d + %d = %d", n1, n2, n1 + n2)
}
```

Agora vamos criar uma calculadora para exemplificar:

```go
package main

import (
    "fmt"
)

func main(){

    var numero1, numero2 int
    var operacao string
    fmt.Print("Digite o primeiro numéro: ")
    fmt.Scanf("%d", &numero1)
    fmt.Print("Digite a operação (+, -, *, /, $): ")
    fmt.Scanf("%d", &operacao)
    fmt.Print("Digite o segundo numéro: ")
    fmt.Scanf("%d", &numero2)

    if operacao == "+" {
        somar(numero1, numero2)
    } else if operacao == "-" {
        resultado := subtrair(numero1, numero2)
        fmt.Printf("%d + %d = %d", numero1, numero2, resultado)
    } else if operacao == "*" {
        resultado := multiplicar(numero1, numero2)
        fmt.Printf("%d x %d = %d", numero1, numero2, resultado)
    } else if operacao == "/" {
        resultado, resto := dividir(numero1, numero2)
        fmt.Printf("%d / %d = %d", numero1, numero2, resultado)
        fmt.Printf("%d %% %d = %d", numero1, numero2, resto)
    }
}

func somar(n1 int, n2 int) {
    fmt.Printf("%d + %d = %d", n1, n2, n1 + n2)
}

func subtrair(n1 int, n2 int) int {
    return n1 - n2
}

// Aqui já podemos definir a variavel que será retornada
func multiplicar(n1 int, n2 int) (resultado int) {
    resultado = n1 * n2
    return
}

// Retornando multiplos valores
func dividir(n1 int, n2 int)(int, int){
    quociente := n1 / n2
    resto := n1 % n2
    return quociente, resto
}
```

