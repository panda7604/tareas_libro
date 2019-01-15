// segundo ejercicio y primero de los de echo, realizado en clase
package main

import (
	"fmt"
	"os" // ayuda a pasar argumentos permite leer esos argumentos
)

func main() { // las llaves deben ir en su lugar
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = "   --  " //esos guiones los puse de prueba se pueden omitir
	}
	fmt.Println(s)
}
