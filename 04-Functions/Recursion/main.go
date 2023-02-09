package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func main() {
	total := factorial(3)
	fmt.Println("El factorial es: ", total)
}
