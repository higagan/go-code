package main

import (
	"fmt"
)

func main() {

	square := num_square(8)
	fmt.Printf("%v", square)
}
func num_square(n int) map[int]int {

	my_map := make(map[int]int)
	for i := 1; i <= n; i++ {
		my_map[i] = i * i
	}
	return my_map
}
