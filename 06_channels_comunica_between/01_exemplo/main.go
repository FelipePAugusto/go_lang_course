package main

import "fmt"

// Thread 1
func main() {

	// Meio de comunicação entre a Thread 1 e a Thread 2
	hello := make(chan string) // Declaração de um chanel do tipo string

	// Thread 2
	go func() { // Routine ou MultiThread que irá atribuir um valor a uma variavel que está na Thread 1 por meio do chanel hello
		hello <- "hello world!"
	}()

	result := <-hello // Toda vez que tiver um valor neste canal(hello) manda ele pra variavel result

	fmt.Println(result)
}
