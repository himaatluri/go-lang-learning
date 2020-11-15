package main

import "fmt"

var color string = "Red"
var cards = []string{"Spade", "Diamond"}

func main() {
	fmt.Println(color)
	cards = append(cards, "Hearts")
	for index, card := range cards {
		fmt.Println(index, card)
	}
}
