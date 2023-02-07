package main

import "fmt"

func main() {
	var name string
	var lastname string
	edad := 30
	name = "Francisco"
	lastname = "Molina"

	fmt.Println(name, lastname, edad)

	//Los valores iniciales de las variables no son como en js: a es 0, b es 0, c es un espacio y el bool es false/true
	var a int
	var b float32
	var c string
	var d bool

	fmt.Println(a, b, c, d)
}
