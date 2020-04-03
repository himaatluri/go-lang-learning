package main

import "fmt"

func main() {
	// creating a var and assigning a value

	//var card string = "Ace of Spades"
	// long form method:
	//
	// var:             We are about to create a new variable, short for variable
	// card: 			The name of the variable we will be greeting
	// string: 			Only a string will ever be assigned to this variable
	// "Ace of Spades": Assign the value "Ace of Spades" to this variable

	// short form mode
	// the Go compiler will figure out the data type
	// := this will let the Go complier to define the type
	// we only use this on the very initial definition of a var
	// card := "Ace of Spades"
	card := newCard()

	// if i have to re assign a new value later
	// no need to put ':=', doing this will throw error
	// card = "Five of Diamonds"

	fmt.Println(card)
}

// new function

// string: is needed to tell what data type we are going to return
func newCard() string {
	// return the statement when this func is called
	return "Five of Diamonds"
}
