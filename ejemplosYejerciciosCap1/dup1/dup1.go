package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++

	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("la palabra %s la repetiste %d veces \n", line, n)
		}
	}

}

/* ojo para probar el ejercicio a mi me funciono desde el cmd
se corre asi
go run dup1.go
hola
hola
perro
perro
12345
adios
en fin las palabras que gustes contara las repetidas. enseguida le das ctrl+c
*/
