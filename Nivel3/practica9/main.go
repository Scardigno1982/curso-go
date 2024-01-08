package main

import "fmt"

func main() {
	// Definir una variable de tipo string
	deporteFav := "pin pong"

	// Usar un switch statement con la expresión de switch especificada como la variable deporteFav
	switch deporteFav {
	case "fútbol":
		fmt.Println("El deporte favorito es fútbol")
	case "baloncesto":
		fmt.Println("El deporte favorito es baloncesto")
	default:
		fmt.Println("El deporte favorito no es ni fútbol ni baloncesto")
	}
}
