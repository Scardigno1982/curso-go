package main

import "fmt"

func main() {
	// Crear una variable de tipo string usando un string literal no interpretado
	str := `Hello, world! This is a raw string literal. It can include "quotes" and
newlines.`

	// Imprimir la variable
	fmt.Println(str)
}
