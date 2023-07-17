package main

import "fmt"

// Estructura de Datos
type car struct {
	brand string
	year  int
	color string
	seat  int
}

func main() {
	mycar := car{brand: "Ford", year: 2020, color: "Red", seat: 4}
	fmt.Println(mycar)

	//Otra manera
	var otherCar car
	otherCar.brand = "Ferrari"
	otherCar.color = "black"
	fmt.Println(otherCar)
}
