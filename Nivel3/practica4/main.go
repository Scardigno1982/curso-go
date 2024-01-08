package main

import "fmt"

var j int = 0

func main() {

	anio := 1982

	for {
		if anio > 2024 {
			break
		}
		fmt.Println(anio)
		anio++
		j++
	}

	fmt.Println("Total de años: ", j)
}
