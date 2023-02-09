package main

import (
	"fmt"
)

func main() {

	if nombre, edad := "Fran", 26; nombre == "Francisco" {
		fmt.Println("Hola,", nombre)
	} else {
		fmt.Printf("Nombre: %s - Edad: %d\n", nombre, edad)
	}

	//Obtener valor de mapa
	users := make(map[string]string)

	users["Francisco"] = "Francisco@gmail.com"
	users["Molina"] = "Molina@gmail.com"

	//correo, ok := users["Francis"]

	if correo, ok := users["Molina"]; ok {
		fmt.Println(correo)
	} else {
		fmt.Println("No fue posible obtener el valor")
	}

}
