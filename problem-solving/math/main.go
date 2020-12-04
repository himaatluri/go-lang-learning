package main

import (
	"fmt"
)

func addNumbers(a int, b int) int {
	var c = a + b
	return c
}

func main() {
	var a, b, c int

	fmt.Println("Enter the first number: ")
	fmt.Scan(&a)

	fmt.Println("Enter the second number: ")
	fmt.Scan(&b)

	c = addNumbers(a, b)

	fmt.Println("Output: ")
	fmt.Println("Sum: ", c)
}
