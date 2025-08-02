package main

import (
	"fmt"
	"lesson-2/tempconv"
)

func main() {
	z := tempconv.AbsolutZeroC
	fmt.Printf("z = %v, %.2f fahrenheit\n", z, tempconv.CToF(z))
	fmt.Println(tempconv.Smt)
}
