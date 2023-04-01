package main

import "fmt"

func main() {

	//Declaración de constantes
	const pi float64 = 3.14 // Nunca cambiará de valor ene el tiempo
	const pi2 = 3.1415

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	//Declaración de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Area  cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Área cuadrado:", areaCuadrado)

	//Operaciones Matemáticas

	x := 10
	y := 50

	//Suma
	result := x + y
	fmt.Println("Suma:", result)

	//Resta
	result = y - x
	fmt.Println("Resta:", result)

	//Multiplicación
	result = x * y
	fmt.Println("Multiplicación:", result)

	//Divición
	result = y / x
	fmt.Println("División:", result)

	//Módulo o residuo - común para saber si un número es par o impar
	result = y % x
	fmt.Println("Residuo:", result)

	//Incremental = x + 1
	x++
	fmt.Println("Incremental:", x)

	//Decremental
	x--
	fmt.Println("Decremental:", x)

	//Retos
	// Hallar el área de un rectángulo, trapecio y círculo

	//Área Rectángulo A= altura * base
	altura = 30
	base = 50
	areaRectangulo := altura * base
	fmt.Println("Área Rectángulo:", areaRectangulo)

	//Área de trapecio A = altura * ((lado 1 + lado 2)/2)

	lado1 := 40
	lado2 := 55
	var areaTrapecio float64 = float64(altura) * ((float64(lado1) + float64(lado2)) / 2)
	fmt.Println("Área Trapecio:", areaTrapecio)

	//Área de círculo  A = π * (radio * radio)

	var π float64 = 3.14159265
	var radio float64 = 3.5
	areaCirculo := π * (radio * radio)
	fmt.Println("Área círculo:", areaCirculo)

}
