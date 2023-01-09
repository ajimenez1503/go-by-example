package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.2", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("12", 0, 64)
	fmt.Println(i)

	k, _ := strconv.Atoi("123")
	fmt.Println(k)

	_, e := strconv.Atoi("dsfas")
	fmt.Println(e)
}
