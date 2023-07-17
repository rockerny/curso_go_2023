package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {
	var textMinuscula = strings.ToLower(text)
	fmt.Println(textMinuscula)
	var textReverse string

	for i := len(textMinuscula) - 1; i >= 0; i-- {
		textReverse += string(textMinuscula[i])

	}

	if textMinuscula == textReverse {
		fmt.Println("Es un Palindromo")
	} else {
		fmt.Println("No es un Palindromo")
	}
}

func main() {
	slice := []string{"hola", "que", "hace"}

	for i, valor := range slice {
		fmt.Println(i, valor)
	}

	for _, valor := range slice {
		fmt.Println(valor)
	}

	for i := range slice {
		fmt.Println(i)
	}

	isPalindromo("Amor a Roma")

}
