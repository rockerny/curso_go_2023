package main

import "fmt"

func main() {

	//Declaración de variables
	helloMessage := "Hello"
	worldMessage := "world"

	//Println -- Agrega un salto de líneal despues de imprimir
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	//Print.f -- Agrega un valor a una variable string
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos) // %s valor númerico, %d valor string
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos) // \%v cuando no conocemos el valor dela variable

	//Sprintf -- Todo se guarda en una nueva variable
	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos)
	fmt.Println(message)

	//Tipo de datos -- Con Printf podemos saber que tipo de datos tiene una variable
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)
}
