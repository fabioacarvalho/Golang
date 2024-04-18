# Golang
 
## Summary

- Basic
    - 1 - Preparação do Ambiente e Conceitos Iniciais
    - 2 - Variáveis e Tipos de Dados
    - 3 - Operadores Aritméticos
    - 4 - Estruturas Condicionais
    - 5 - Estruturas de repetição

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

