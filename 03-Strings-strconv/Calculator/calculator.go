package main

import (
	"fmt"
	"strconv"
	"strings"
)

func add(expression string) int {

	values := strings.Split(expression, "+")
	var result int

	for _, value := range values {
		num, error := strconv.Atoi(value)

		if error != nil {
			fmt.Println(error)
			fmt.Println("Solo puede resolver sumas con el operador + ")
		} else {
			result += num
		}

	}
	return result
}

func main() {

	var expression string
	var result int

	fmt.Print("=>")

	fmt.Scanln(&expression)

	result = add(expression)

	fmt.Printf("==> %d \n", result)

}
