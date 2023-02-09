package main

import "fmt"

func main() {
	nombres := [...]string{"Francisco", "Moli", "Francisco Molina"}

	//For comun:
	for i := 0; i < len(nombres); i++ {
		fmt.Println(nombres[i])
	}

	//For each:
	for indice, elemento := range nombres {
		fmt.Println(indice, elemento)
	}
}
