package main

import "fmt"

func main() {
	// Crear una constante con tipo
	const myTypedConst int = 10
	fmt.Println("Constante con tipo:", myTypedConst)

	// Crear una constante sin tipo
	const myUntypedConst = "Hello, world!"
	fmt.Println("Constante sin tipo:", myUntypedConst)
}
