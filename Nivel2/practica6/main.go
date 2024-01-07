package main

import "fmt"

const (
	// Usar iota para generar constantes para los próximos 4 años
	year1 = 2022 + iota
	year2
	year3
	year4
)

func main() {
	// Imprimir los valores de las constantes
	fmt.Println(year1, year2, year3, year4)
}
