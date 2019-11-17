package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(value int) int {
		sum += value
		return sum
	}
}

func main() {
	f := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0 + 1 + .. + %d = %d\n", i, f(i))
	}
}
