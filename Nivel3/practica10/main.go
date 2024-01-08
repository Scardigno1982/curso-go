package main

import "fmt"

func main() {
	fmt.Println(true && true)  // Imprime: true
	fmt.Println(true && false) // Imprime: false
	fmt.Println(true || true)  // Imprime: true
	fmt.Println(true || false) // Imprime: true
	fmt.Println(!true)

}
