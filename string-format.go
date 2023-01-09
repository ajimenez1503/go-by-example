package main

import "fmt"

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	fmt.Printf("struct: %v\n", p)
	fmt.Printf("struct: %+v\n", p)
	fmt.Printf("struct: %#v\n", p)

	fmt.Printf("type: %T\n", p)

	fmt.Printf("bool: %t\n", true)

	fmt.Printf("int: %d\n", 123)

	fmt.Printf("binary: %b\n", 134)

	fmt.Printf("char: %c\n", 33)

	fmt.Printf("hex: %x\n", 456)

	fmt.Printf("float: %f\n", 78.7)

	fmt.Printf("science: %e\n", 12300000)

	fmt.Printf("str: %s\n", "string")
	fmt.Printf("str-quate: %s\n", "string")

	fmt.Printf("Pointer: %p\n", &p)

	fmt.Printf("width: |%6d|\n", 12)
	fmt.Printf("width: |%6.3f|\n", 12.1)
	fmt.Printf("width: |%6s|\n", "test")

	s := fmt.Sprintf("test %s", "string")
	fmt.Println(s)

}
