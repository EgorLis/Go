package main

import (
	"examples/Lesson-2/tempconv"
	"fmt"
)

func main() {
	z := tempconv.AbsolutZeroC
	fmt.Printf("z = %v, %.2f fahrenheit\n", z, tempconv.CToF(z))
	fmt.Println(tempconv.Smt)
}
