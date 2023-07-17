package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)

	m["José"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	// Recorrer map

	for i, v := range m {
		fmt.Println(i, v)
	}

	//Encontrar valor
	value, ok := m["José"]
	fmt.Println(value, ok)
}
