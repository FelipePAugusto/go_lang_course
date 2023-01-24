package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type ClienteNome string
type ClienteEmail string
type ClienteCPF string

type Cliente struct {
	Nome  string
	Email string
	CPF   int
}

func (c Cliente) Exibe() {
	fmt.Println("Exibindo Cliente pelo metodo: ", c.Nome)
}

type ClienteInternacional struct {
	Cliente
	Pais string `json:"pais"` // a crase (ou BACKTICK) nesse ponto serve para dizer que quando for JSON, será com essa nomenclatura
}

func main() {
	// var ClienteNome clientenome

	cliente := Cliente{
		Nome:  "Felipe",
		Email: "teste@email.com",
		CPF:   12345678,
	}

	fmt.Println(cliente)

	cliente2 := Cliente{"Joao", "valida@email.com", 123987748}

	fmt.Println(cliente2)

	// cliente3 := ClienteInternacional{
	// 	Nome:  "Joao",
	// 	Email: "valida@email.com",
	// 	CPF:   87128731,
	// 	Pais:  "Japao",
	// }

	// fmt.Println(cliente3)

	cliente4 := ClienteInternacional{
		Cliente: Cliente{
			Nome:  "Marcos",
			Email: "marcos@email.com",
			CPF:   246351,
		},
		Pais: "Africa do Sul",
	}

	fmt.Println(cliente4)
	fmt.Printf("Nome: %s - Email: %s\n", cliente4.Nome, cliente4.Cliente.Email)

	// METODO dentro de STRUCT
	cliente.Exibe()
	cliente2.Exibe()
	cliente4.Exibe()

	// JSON

	clienteJson, err := json.Marshal(cliente4)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(clienteJson) // Pega os dados em um slice de bites
	fmt.Println(string(clienteJson))

	jsoncliente4 := `{"Nome":"Marcos","Email":"marcos@email.com","CPF":246351,"pais":"Africa do Sul"}`
	cliente5 := ClienteInternacional{}

	json.Unmarshal([]byte(jsoncliente4), &cliente5) // Ele usando referencia(espaço de memoria) faz com que altere o conteudo da variavel ClienteInternacional q é uma struct
	fmt.Println(cliente4)
}
