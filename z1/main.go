package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var name string
	name = "World"

	var i int
	var j float64
	var b bool

	os.Stderr.WriteString(fmt.Sprintf("Hello %s %d %f %t\n", name, i, j, b))
	os.Stderr.WriteString(fmt.Sprintf("%.10f\n", math.Pi))
}
