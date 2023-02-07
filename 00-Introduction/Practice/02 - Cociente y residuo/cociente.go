/*
Practica 02: Calcular cociente y el Residuo de dos números enteros
Enunciado: halar el cociente y el residuo (resto) de dos números enteros.

Análisis: para la solución de este problema, se requiere que el usuario ingrese dos números enteros y el sistema realice el cálculo respectivo para hallar el cociente y residuo.

*/

package main

import "fmt"

//función que calcula cociente
func cociente(a, b int) int {
	return a / b
}

//Función que calcula residuo
func residuo(a, b int) int {
	return a % b
}

//Función principal
func main() {

	//Variables
	var a, b, c, r int

	//Entrada  de datos
	fmt.Print("Ingrese N01: ")
	fmt.Scanln(&a)

	fmt.Print("Ingrese N02: ")
	fmt.Scanln(&b)

	//Llamar la función cociente
	c = cociente(a, b)

	//LLamar la función residuo
	r = residuo(a, b)

	//Salida de datos
	fmt.Println("Cociente:", c)
	fmt.Println("Residuo:", r)
}
