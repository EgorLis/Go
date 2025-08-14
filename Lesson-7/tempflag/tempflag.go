package main

import (
	"examples/Lesson-2/tempconv"
	"flag"
	"fmt"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "температура")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
