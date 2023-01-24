package main

import "fmt"

var u = 15 // var can created anywhere but in the escope CodeBlocks

func main() {
	nr_bytes, erros := fmt.Println("Hello, world!", "Olá", "Welcome", 100) //How many bytes it used to print and if there is any error
	// I can put how many stuffs as i want and any kind of parameters
	fmt.Println(nr_bytes, erros)

	// I can put  _ in places that i dont want use the variable
	_, variavel_de_erro := fmt.Println("Hellzo, world!", "TEST", 35)
	fmt.Println(variavel_de_erro)

	// The Marmote or declaration, already define the type of variables
	x := 10
	y := "string"

	fmt.Printf("x: %v, %T \n", x, x)
	fmt.Printf("y: %v, %T \n\n", y, y)

	x, z := 20, 30

	fmt.Printf("x: %v, %T \n", x, x)
	fmt.Printf("z: %v, %T \n", z, z)

	// THere are escope an CodeBlocks, the declaration of variables working only inside the escope, not global
	// var teste = 123

	// TERMINOLOGIA
	// KeyWords, words reserved for only functions especific
	// break        default      func         interface    select
	// case         defer        go           map          struct
	// chan         else         goto         package      switch
	// const        fallthrough  if           range        type
	// continue     for          import       return       var

	// Statements (Declaration of variables or functions)

	//ERRO -> u = 15.1 // Types declarations will be ever same type until final

	// Every time that declaration one variable and don't inicialize it, it will be 0 or 0 or false or " "

	// We got it defining a new type declaration, for example hotdog of type int -> type hotdog int
	// TYPE CONVERSION or TYPE CASTING or TYPE COERCION
	// x = int(hotdog)
}
