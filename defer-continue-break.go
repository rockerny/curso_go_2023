package main

import "fmt"

func main() {

	//Defer
	defer fmt.Println("Hola") //Ejecuta al final de la función
	fmt.Println("Mundo")

	//Continue y break
	for i := 0; i < 10; i++ {
		//Continue
		if i == 2 {
			fmt.Println("Es 2")
			continue //Corta la acción de for y continua
		}

		//Break
		if i == 8 {
			fmt.Println("Break")
			break
		}
		fmt.Println(i)

	}

}
