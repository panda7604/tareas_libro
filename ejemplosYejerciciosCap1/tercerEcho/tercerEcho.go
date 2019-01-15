// cuarto ejercicio y tercer echo realizado en clase
package main

import (
	"fmt"
	"os" // ayuda a pasar argumentos permite leer esos argumentos
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], "   --   "))
}
