package main

import "fmt"

func main() {
	// Imprimir todos los números del 1 al 10,000
	for i := 60; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("%#U\n", i)
		}
	}
}
