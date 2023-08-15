package main

import (
	//pk es el alias
	pk "curso_golang_platzi/src/mypackage"
	"fmt"
)

func main() {
	a := 50
	//apunta a la dirección de memoria de las variable a
	b := &a

	fmt.Println(b)
	//* Apunta al valor que esta en la variable de memoria
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

	myPC := pk.Pc{Ram: 16, Disk: 200, Brand: "msi"}
	fmt.Println(myPC)

	myPC.Ping()

	fmt.Println(myPC)
	myPC.DuplicateRAM()

	fmt.Println(myPC)
	myPC.DuplicateRAM()

	fmt.Println(myPC)
	myPC.DuplicateRAM()

	fmt.Println(myPC)

	//Stringers: Personalizar el output de Struct
}



//mypackage.go


package mypackage

import "fmt"

// Se usa mayúsculas para indicar que es pública PC con accesos público
type Pc struct {
	Ram   int
	Disk  int
	Brand string
}

func (myPC Pc) Ping() {
	fmt.Println(myPC.Brand, "Pong")

}

func (myPC *Pc) DuplicateRAM() {
	myPC.Ram = myPC.Ram * 2
}

func (myPC Pc) String() string {

	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco y es una %s", myPC.Ram, myPC.Disk, myPC.Brand)
}
