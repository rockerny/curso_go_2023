package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	valor1 := 0
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Print("No es 1\n")
	}

	// With And
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es verdad")
	} else {
		fmt.Print("Es falso\n")
	}

	//With Or
	if valor1 == 0 || valor2 == 3 {
		fmt.Println("Es verdad, OR")
	}

	//Convertir texto anúmero
	value, err := strconv.Atoi("67")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value:", value)

	// Identificar si es Par o Impar

	if parNumero(34) {
		fmt.Println("El número ingresado es Par")
	} else {
		fmt.Println("El número ingresado es Impar")
	}

	// Validación de Password
	if userPassword("Rockerny", "Rocky2345!") {
		fmt.Println("Acceso correcto")
	} else {
		fmt.Println("Acceso Denegado")
	}

}

func parNumero(n int) bool {
	return n%2 == 0

}

func userPassword(u, p string) bool {
	return u == "Rockerny" && p == "Rocky2345!"
}
