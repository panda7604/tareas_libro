package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"comida para pancho": 50,
		"cacahuates": 15,
		"galletas":   20,
		"cocacola":   14,
		"socks":      15}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
