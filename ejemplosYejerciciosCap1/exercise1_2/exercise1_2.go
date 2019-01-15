// Exercis e 1.2: Modify the echo program to print the index and value ofeach of its arguments, one perline.
package main

import (
	"fmt"
	"os"
)

func main() {

	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = "   --  " //esos guiones los puse de prueba se pueden omitir
		//fmt.Println(os.Args[i], " (es el index) ", i)
		fmt.Println("el index ", i, " tiene el valor:...", os.Args[i], "y esa palabra vale", len(os.Args[i]), "caracteres!!")
	}

}
