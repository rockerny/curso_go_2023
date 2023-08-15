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
