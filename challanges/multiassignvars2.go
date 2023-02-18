package main

import "fmt"

func main() {
	var (
		planet string
		isTrue bool
		temp   float64
	)

	planet, isTrue, temp = "Mars", true, 19.5

	println("Air is good on ", planet)
	println("It's ", isTrue)
	fmt.Println("It is", temp, "degrees")
}
