package main

import "fmt"

func main() {
	numbers := make([]int, 0, 5) //Es longitud 0

	addNumbers := append(numbers, 100)

	fmt.Println(numbers, len(numbers), cap(numbers)) //[] 0 5

	fmt.Println(addNumbers, len(addNumbers), cap(addNumbers)) //[100] 1 5

	moreNumbers := make([]int, 2, 3) //Es longitud 2

	moreNumbers[0] = 3
	moreNumbers[1] = 200

	fmt.Println(moreNumbers, len(moreNumbers), cap(moreNumbers)) //[3 200] 2 3
}
