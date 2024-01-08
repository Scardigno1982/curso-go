package main

import "fmt"

var anio int = 1982
var f int = 0

func main() {
	// Imprimir todos los números del 1 al 10,000
	for i := anio; i <= 2024; i++ {
		fmt.Println(i)
		f++
	}

	fmt.Println("Total de años: ", f)
}
