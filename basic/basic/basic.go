package main

import (
	"fmt"
	"math"
)

func add(a, b int) int {
	return a + b
}

func triangle() {
	a, b := 3, 4
	fmt.Println(caTriangle(a, b))
}

func caTriangle(a, b int) int {
	return int(math.Sqrt(float64(a*a + b*b)))
}

func main() {
	fmt.Println(add(1, 2))
}
