package main

import "fmt"

func main() {
	a := 21
	b := 30
	result := a + b

	fmt.Println("El resultado de la suma es: ", result)

	result = b - a
	fmt.Println("El resultado de la resta es: ", result)

	result = a * b
	fmt.Println("El resultado de la multiplicación es: ", result)

	var divi float32 = 3.0 / 2.0
	fmt.Println("El resultado de la division es: ", divi)

	result = a % b
	fmt.Println("El modulo es: ", result)
}
