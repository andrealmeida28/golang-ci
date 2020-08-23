package main

import (
	"fmt"
)

func main() {
	x := 5
	y := 5
	resultado := Soma(x, y)
	fmt.Printf("Resultado da soma (%d+%d) é %d", x, y, resultado)
}

// Soma returns x + y
func Soma(x, y int) int {
	return x + y
}
