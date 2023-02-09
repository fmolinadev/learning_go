package main

import "fmt"

func main() {
	a := 2
	b := 5
	var result bool

	//Se compara la igualdad con dos igual "==":

	result = a == b
	fmt.Printf("%d y %d son iguales?: %t \n", a, b, result) //false

	//Se compara la desigualdad con "!=":

	result = a != b
	fmt.Printf("%d y %d son distintos?: %t \n", a, b, result) //true

	//Se compara el mayor ">":

	result = a > b
	fmt.Printf("%d es mayor que %d?: %t \n", a, b, result) //false

	//Se compara el mayor o igual">=":

	result = a >= b
	fmt.Printf("%d es mayor o igual que %d?: %t \n", a, b, result) //false

	//Se compara el menor "<":

	result = a < b
	fmt.Printf("%d es menor que %d?: %t \n", a, b, result) //true

	//Se compara el menor o igual"<=":

	result = a <= b
	fmt.Printf("%d es menor o igual que %d?: %t \n", a, b, result) //true
}
