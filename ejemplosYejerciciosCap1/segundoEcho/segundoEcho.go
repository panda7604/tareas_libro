// tercer ejercicio y segundo de echo realizado en clase
package main

import (
	"fmt"
	"os" // ayuda a pasar argumentos permite leer esos argumentos
)

func main() {
	s, sep := " ", " "
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = "   --  "
	}
	fmt.Println(s)
}
