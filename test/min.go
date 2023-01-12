package main

import (
	"fmt"
)

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Printf("IntMin(2, -2) = %d, expected -2 \n", IntMin(2, -2))
}
