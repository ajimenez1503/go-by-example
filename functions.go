package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusplus(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) {
	return 5, 6
}

func sum(nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(nums, total)
}

func main() {
	res := plus(1, 2)
	fmt.Println("1 + 2=", res)

	res = plusplus(1, 2, 3)
	fmt.Println("1 + 2 + 3 = ", res)

	a, b := vals()
	fmt.Println(a, b)

	_, c := vals()
	fmt.Println(c)

	sum(1, 2)
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
