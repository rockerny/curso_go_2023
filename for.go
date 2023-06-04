package main

import "fmt"

func main() {
	//For Condicional
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("\n")

	//For while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	//For forever
	/*
		counterForever := 0
		for {
			fmt.Println(counterForever)
			counterForever++
		}*/

	//For inverso

	fmt.Println("\n")

	counterFinal := 10

	for counterFinal > -1 {
		fmt.Println(counterFinal)
		counterFinal--
	}

}
