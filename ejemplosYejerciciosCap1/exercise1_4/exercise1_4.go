//este es el ejercicio 1.4 el cual debe de mostrar lo que se repite y en que archivo se repite
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	fileCount := make(map[string]map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {

		countLines(os.Stdin, counts, fileCount)

	} else {

		for _, arg := range files {

			f, err := os.Open(arg)

			if err != nil {

				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)

				continue

			}

			countLines(f, counts, fileCount)

			f.Close()

		}

	}

	for line, n := range counts {

		if n > 1 {

			fmt.Printf("%d\t%s", n, line)

			for fn, fc := range fileCount[line] {

				fmt.Printf("\n\t -> %dveces en el archivo %s", fc, fn)

			}

			fmt.Println()

		}

	}

}

func countLines(f *os.File, counts map[string]int, fileCount map[string]map[string]int) {

	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		fileName := f.Name()
		counts[text]++

		if fileCount[text] == nil {

			fileCount[text] = make(map[string]int)

		}
		fileCount[text][fileName]++
	}

}
