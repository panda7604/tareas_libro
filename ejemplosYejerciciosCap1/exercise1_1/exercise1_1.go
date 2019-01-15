//Exercis e 1.1: Modify the echo program to also print os.Args[0],the name of the command that invoked it.
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ { // se ha modificado para que nos de desde el cero.
		s += sep + os.Args[i]
		sep = "   --  " //esos guiones los puse de prueba se pueden omitir
	}
	fmt.Println(s)
}
