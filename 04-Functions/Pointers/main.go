package main

import "fmt"

//Punteros
func editValue(text *string) {
	fmt.Printf("%p\n", text)
	*text = "Hola desde la funcion"

}

func main() {
	text := "Hola Mundo de Go"
	fmt.Printf("%p\n", &text)
	fmt.Println("Antes de la función:", text)

	editValue(&text) //Referencia

	fmt.Println("Después de la función:", text)
}
