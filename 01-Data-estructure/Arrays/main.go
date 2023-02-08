package main

import "fmt"

func main() {
	//Los arrays se definen con la cantidad de datos y el tipo que va a recibir:
	var allNumbers [10]int
	fmt.Println(allNumbers) //Esto inicializa como: [0 0 0 0 0 0 0 0 0 0]

	allNumbers[2] = 2
	allNumbers[1] = 1
	allNumbers[3] = 3
	fmt.Println(allNumbers) //Esto ahora es: [0 1 2 3 0 0 0 0 0 0]

	//Arrays con valores y longitud definida:

	allNames := [3]string{"Francisco", "Daniel", "Molina"}
	fmt.Println(allNames)

	//Arrays con valores y sin longitud definida:
	allColors := [...]string{"Rojo", "Verde", "Amarillo"}
	fmt.Println(allColors)
	//Para saber la longitud se usa "len":
	fmt.Println(len(allColors))

	//Arrays con valores definidos en indice y sin longitud definida:
	allMoney := [...]string{0: "Dolar", 1: "Euro", 3: "Peso"}
	fmt.Println(allMoney) //Esto ahora es: ["Dollar" "Euro" "un espacio vacio" "Pes0" ] Los valores strings no definidos son un espacio.

	//Se pueden crear sub-arrays:
	subMoney := allMoney[0:3]
	fmt.Println(subMoney) //Esto imprime desde el indice 0 hasta el 2, lo limita el 3: => allMoney[0:3]
}
