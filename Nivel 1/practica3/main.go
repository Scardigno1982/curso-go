package main

import "fmt"

// Declaración de tres variables a nivel de paquete
var x int
var y string
var z bool

func main() {
	x = 42
	y = "James Bond"
	z = true

	// Usa fmt.Sprintf para imprimir todos los valores en un solo string
	// Asigna el valor retornado de tipo string a la variable "s"
	s := fmt.Sprintf("%v %v %v", x, y, z)

	// Imprime el valor almacenado por la variable "s"
	fmt.Println(s)
}
