package main

import "fmt"

func main() {

	//La diferencia es que un slicen es mutable:

	allLetters := []string{"a", "b", "c"}

	fmt.Println(allLetters, len(allLetters)) //[a b c] 3

	//Para agregar datos se usa "append". Recibe como primer parametro el slicen y segundo lo que va a agregar:

	addLetters := append(allLetters, "d")

	fmt.Println(addLetters, len(addLetters)) //[a b c d] 4

	//Obtener un fragmento en un subSlicen:

	subLetters := addLetters[:2]

	fmt.Println(subLetters) //Esto me devuelve: [a b]

	allLetters[0] = "z"

	fmt.Println(allLetters)

	/* Punteros, logitud y capacidad */

	meses := []string{"Enero", "Febrero", "Marzo", "Abril"}

	fmt.Printf("Len: %v, Cap: %v, Ref: %p \n", len(meses), cap(meses), meses)

	//Punteros:

	//Longitud:

	//Capacidad:

}
