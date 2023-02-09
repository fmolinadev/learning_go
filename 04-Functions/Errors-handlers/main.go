package main

import (
	"errors"
	"fmt"
)

func part(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("No es posiible dividir entre 0!!")
	} else {
		return dividendo / divisor, nil
	}

}

func main() {

	result, error := part(4, 2)

	if error == nil {
		fmt.Println("Resultado: ", result)
	} else {
		fmt.Println(error)
	}

}
