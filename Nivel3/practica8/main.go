package main

import "fmt"

func main() {
	// Definir una variable
	num := 0

	// Usar un switch statement sin expresión especificada
	switch {
	case num > 5:
		fmt.Println("El número es mayor que 5")
	case num == 0:
		fmt.Println("El número es cero")
	case num < 5:
		fmt.Println("El número no es mayor que 5")
	}
}
