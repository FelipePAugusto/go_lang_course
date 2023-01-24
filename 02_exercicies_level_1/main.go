package main

import "fmt"

func main() {
	x, y, z := 42, "Jack", true

	fmt.Printf("%v %T %v %T %v %T\n\n", x, x, y, y, z, z)

	fmt.Printf("%v, %T\n\n", x, x)
	fmt.Printf("%v, %T\n\n", y, y)
	fmt.Printf("%v, %T", z, z)
}
