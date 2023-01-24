package main

import (
	"fmt"
	"runtime"
	"time"
)

func contador(tipo string) {
	for i := 0; i < 5; i++ {
		fmt.Println(tipo, i)
		time.Sleep(time.Second)

	}
}

func main() {
	go contador("sem go routine")
	go contador("com go routine")

	fmt.Println("Hello, world!")
	fmt.Println("Hello, space!")

	fmt.Println("end...")
	time.Sleep(time.Second * 10)

	runtime.GOMAXPROCS(1) // Estipula o numero maximo de CPUs a serem utilizadas para execução do programa
	fmt.Println("Começou...")

	go func() { // GoRoutine
		for { // FOR sem condição, é um loop infinito

		}
	}()

	time.Sleep(time.Second)
	fmt.Println("Terminou...")
}
