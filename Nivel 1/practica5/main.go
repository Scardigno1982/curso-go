package main

import "fmt"

// Crear mi propio tipo
type hotdog int

var x hotdog
var y int

func main() {

	fmt.Println(x)
	fmt.Printf("%T \n", x)

	x = 12
	fmt.Println(x)

	y = int(x)

	// Imprime el valor almacenado en “y”
	fmt.Println(y)
	// Imprime el tipo de “y”
	fmt.Printf("%T\n", y)
}
