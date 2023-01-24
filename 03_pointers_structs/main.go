package main

import "fmt"

type Carro struct {
	Name string
}

func (c *Carro) andar() {
	c.Name = "Renault"
	fmt.Println(c.Name, " Andou")
}

func abc(a *int) { // Recebido como parametro uma variavel do tipo ponteiro
	*a = 200 // Faz a atribuicao do valor usando ponteiro
}

func main() {
	a := 10 // Variavel A esta em um espaco de memoria

	var ponteiro *int = &a // Variavel ponteiro foi declarada com tipo ponteiro e inteiro e foi atribuido o valor de memoria da variavel A

	*ponteiro = 50 // Utilizando o ponteiro foi atribuido o valor de 50 no espaço de memoria da variavel A

	b := &a        // Aqui foi feito a mesma coisa
	*b = 60        // Atribuido valor 60 no espaco de memoria da variavel A
	fmt.Println(a) // 60

	variavel := 10 // Declarado variavel e atribuido com valor 10

	abc(&variavel) // Chamado a funcao passando o espaco de memoria como parametro

	fmt.Println(variavel)

	carro := Carro{
		Name: "Fox",
	}

	carro.andar()
	fmt.Println(carro.Name)
}
