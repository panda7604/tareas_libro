package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) // aca ftoc esta recibiendo freezing y freezing vale 32.0
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))
}

func fToC(loquesea float64) float64 {
	return (loquesea - 32) * 5 / 9

}
