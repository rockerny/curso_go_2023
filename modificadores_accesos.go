//export GOROOT=/usr/local/go


package main

import (
	pk "curso_golang_platzi/src/mypackage"
	"fmt"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2020
	fmt.Println(myCar)

	//var myOtherCar pk.carPrivate
	//fmt.Println(myOtherCar)

	pk.PrintMessage("Hola Ada")
}



/////mypackage.go


package mypackage

import "fmt"

// Se usa mayúsculas para indicar que es pública CarPublic con accesos público
type CarPublic struct {
	Brand string
	Year  int
}

type carPrivate struct {
	brand string
	year  int
}

// PrintMessage imprimir un mensaje
func PrintMessage(text string) {
	fmt.Println(text)
}
