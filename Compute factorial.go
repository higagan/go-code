package main

import (
	"fmt"
)

func main() {

	fact := factorial(8)
	fmt.Print(fact)
}
func factorial(n int) int {

	res := 1
	for i := n; i > 0; i-- {
		res *= i
	}
	return res
}
