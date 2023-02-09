package main

import (
	"fmt"
	"strings"
)

func reverseText(letters string) string {
	arrayLetters := strings.Split(letters, "")
	arrayInvert := make([]string, 0)

	for i := len(arrayLetters) - 1; i >= 0; i-- {
		arrayInvert = append(arrayInvert, arrayLetters[i])
	}

	return strings.Join(arrayInvert, "")
}

func isPalindrome(text string) bool {
	text = strings.ToLower(text)
	text = strings.ReplaceAll(text, " ", "")

	textInReverse := reverseText(text)

	return textInReverse == text
}

func main() {
	if isPalindrome("Oso") {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}
