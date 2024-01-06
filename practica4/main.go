package main

import "fmt"

// Crear mi propio tipo
type miTipo int

var x miTipo

func main() {
	// imprime el valor de la variable x
	fmt.Println(x)

	//imprime el tipo de variable
	fmt.Printf("%T \n", x)

	x = 412

	fmt.Println(x)
}
