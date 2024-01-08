package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("Si dividimos %v entre 4, el residuo (o módulo) es %v\n", i, i%4)

	}
}
