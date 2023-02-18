package main

// multiply float with int and re-assign w/ var
import "fmt"

func main() {
	n := 0.
	f := 3.14
	m := 2

	n = f * float64(m)

	fmt.Println(n)
}
