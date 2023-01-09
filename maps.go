package main

import "fmt"

func main() {
	// Declaration
	m := make(map[string]int)

	// Set and get
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map", m)

	v1 := m["k1"]
	fmt.Println("v1", v1)

	// Len
	fmt.Println("len", len(m))

	// Delete
	delete(m, "k2")
	fmt.Println("map", m)

	// Check if exits
	_, prs := m["k2"]
	fmt.Println("present", prs)

	// Declaration and initialization
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map", n)
}
