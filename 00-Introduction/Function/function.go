package main

import "fmt"

func saludar(name string) {
	fmt.Println("Hola!,", name)
}

func sumar(a, b int) int {
	return a + b
}

func main() {
	saludar("Francisco")
	result := sumar(5, 55)
	fmt.Println("La suma es:", result)
}
