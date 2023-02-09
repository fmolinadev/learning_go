package main

import "fmt"

func main() {
	//Bucle tipo while
	/**
	numeros := 12455458
	c := 0
	for numeros > 0 {
		numeros /= 10
		c++
	}
	fmt.Println("Cantidad de digitos es: ", c)
	*/

	//Bucle for

	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	for i := 0; i <= 10; i++ {

		if i == 5 {
			fmt.Println("Salta la iteración")
			continue
		}

		if i == 8 {
			fmt.Println("Romper con bucle ")
			break
		}
		fmt.Println(i)
	}
}
