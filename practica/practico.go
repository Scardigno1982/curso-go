package main

import "fmt"

func main() {
	// Usando el operador de declaración corta, asigna los siguientes valores a variables
	x := 42
	y := "James Bond"
	z := true

	// Imprime los valores almacenados en esas variables usando una sola declaración de la función println
	fmt.Println(x, y, z)

	// Imprime los valores almacenados en esas variables usando múltiples declaraciones de println
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}