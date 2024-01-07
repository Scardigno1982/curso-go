package main

import "fmt"

func main() {
	// Definir algunas variables
	a := 10
	b := 20

	// Escribir expresiones usando los operadores y asignar sus valores a variables
	eq := a < b
	le := a <= b
	ge := a >= b
	ne := a != b
	lt := a < b
	gt := a > b

	// Imprimir los valores de las variables
	fmt.Printf("a < b: %v\n", eq)
	fmt.Printf("a <= b: %v\n", le)
	fmt.Printf("a >= b: %v\n", ge)
	fmt.Printf("a != b: %v\n", ne)
	fmt.Printf("a < b: %v\n", lt)
	fmt.Printf("a > b: %v\n", gt)
}
