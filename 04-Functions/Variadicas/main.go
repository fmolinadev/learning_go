package main

import "fmt"

func add(name string, num ...int) (string, int) {
	var result int
	message := fmt.Sprintf("%s, la suma es: ", name)
	for _, num := range num {
		result += num
	}
	return message, result
}
func main() {
	text, total := add("Fran", 2, 5, 7)
	fmt.Println(text, total)
}
