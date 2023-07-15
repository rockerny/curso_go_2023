package main

import "fmt"

func main() {
	//Array
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array)) //Len muestra cuantos elementos tiene el array y Cap muestra la capacidad máxima del array

	//Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// Métodos en el slice
	fmt.Println(slice[.0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4]) //El primer número es inclusivo y el segundo es exclusivo
	fmt.Println(slice[4:])

	//Append
	slice = append(slice, 7)
	fmt.Println(slice)

	//Append nueva list
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

}
