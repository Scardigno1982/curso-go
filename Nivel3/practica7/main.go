package main

import "fmt"

func main() {
	// Definir una variable
	num := 6
	// Usar un if statement para verificar si el número es mayor que 5
	if num > 5 {
		fmt.Println("El número es mayor que 5")
	} else if num == 0 {
		fmt.Println("El número es cero")
	} else if num < 5 {
		fmt.Println("El número no es mayor que 5")
	}
}
