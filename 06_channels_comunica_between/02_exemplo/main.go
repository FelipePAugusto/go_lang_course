package main

import "fmt"

func main() { // Exemplo de quando o CHANNEL trava por NÃO Receber nenhum valor por comunicação
	forever := make(chan string)

	go func() { // Routine com loop infinito
		x := true

		for {
			if x == true {
				continue
			}
		}
	}()

	fmt.Println("Aguardando para sempre") // Print na tela pra demonstrar que chegou até aqui
	<-forever                             // Caso chegue até aqui no processamento do código e finalize a aplicação, então quer dizer que chegou algum valor na variavel forever
}
