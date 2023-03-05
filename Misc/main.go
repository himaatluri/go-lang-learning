package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, Greetings!")
	ct := time.Now()
	t := []byte("Time: ")
	fmt.Printf(
		"%v\nDate: %v-%v-%v\n",
		string(
			ct.AppendFormat(t, time.Kitchen),
		), ct.Day(), ct.Month(), ct.Year())
}
