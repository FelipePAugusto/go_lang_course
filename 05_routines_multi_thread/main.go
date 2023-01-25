package main

import (
	"fmt"
	"runtime"
	"time"
)

func contador(tipo string) { // tipo de chamada da funcao - com ou sem GO routine
	for i := 0; i < 5; i++ {
		fmt.Println(tipo, i)
		time.Sleep(time.Second)

	}
}

func main() { // Funcao Main é rodada em uma Thread
	go contador("sem go routine") // Routine rodando em paralelismo
	go contador("com go routine") // Routine rodando em uma outra MultiThread

	fmt.Println("Hello, world!")
	fmt.Println("Hello, space!")

	fmt.Println("end...")
	time.Sleep(time.Second * 10)

	runtime.GOMAXPROCS(1) // Estipula o numero maximo de CPUs a serem utilizadas para execução do programa
	fmt.Println("Começou...")

	go func() { // GoRoutine (Isso seria rodado em uma Thread separada)
		for { // FOR sem condição, é um loop infinito

		}
	}()

	time.Sleep(time.Second)
	fmt.Println("Terminou...")
}
