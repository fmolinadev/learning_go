package main

import "fmt"

func main() {
	hello := "Hola"
	world := "Mundo!"

	fmt.Println(hello + " " + world)

	name := "Fran"
	age := 30

	fmt.Printf("Hola %s! Tu edad es %d \n", name, age)

	/**** Si no conozco el tipo de dato puedo usar %v en el fmt
	y si quiero saber el tipo de dato, puedo usar %T ****/

	message := fmt.Sprintf("Hola %s! Esta es la variable y tu edad es %d", name, age)
	fmt.Println(message)

	//Para tomar datos desde el promt se hace asi:

	var lastname string
	fmt.Println("Ingresa tu apellido: ")
	fmt.Scanln(&lastname)
	fmt.Printf("Bienvenido, %s", lastname)

}
