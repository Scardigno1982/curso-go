package main

import "fmt"

func main() {
	// Asignar un int a una variable
	num := 10

	// Imprimir el número en decimal, binario y hexadecimal
	fmt.Printf("Decimal: %d\n", num)
	fmt.Printf("Binario: %b\n", num)
	fmt.Printf("Hexadecimal: %x\n", num)

	// Hacer un shift de bits del número una posición a la izquierda y asignar eso a una variable
	shiftedNum := num << 1

	// Imprimir la variable shift en decimal, binario y hexadecimal
	fmt.Printf("Shifted Decimal: %d\n", shiftedNum)
	fmt.Printf("Shifted Binario: %b\n", shiftedNum)
	fmt.Printf("Shifted Hexadecimal: %x\n", shiftedNum)
}
