package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

var conferenceName string = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50 // cause tickets can't be negative, setting as uint instead of int helps
// var bookings = make([]map[string]string, 0) // defining an array that has max 50 items
var bookings = make([]UserData, 0) // using the struct

// we are creating a custom struct with name UserData
// we can define datatypes of each element
// struct in go == classes in Java
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{} // we create to sync the threads

func main() {
	greetUsers()
	// in Go there is only one loop i.e "for" loop no while B-)

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)

			wg.Add(1) // adding to a wait group
			go sendTicket(userTickets, firstName, lastName, email)
			// the `go` keyword starts a new goroutine, spins a new thread and lets the main code continue
			// this will be a problem if main code exits, we have to synchronize it with main thread
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			// "=" -> for assining values to variables
			// "==" -> for comparing 2 values
			if remainingTickets == 0 {
				// end
				fmt.Println("Our conference is booked out. Come back next year. ")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or Last name you entered is too short")
			}

			if !isValidEmail {
				fmt.Println("Email addres you entered doesn't contain @ sign")
			}

			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}
	}
	wg.Wait() // wait till all group threads done
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings { // _ are used to identify unused vars and this is how we avoid using index in a for_loop
		// firstNames = append(firstNames, booking["firstName"])
		firstNames = append(firstNames, booking.firstName) // just like a class, we obtain values of a struct
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask for user input
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidEmail, isValidName, isValidTicketNumber
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	// user data
	// var userData = make(map[string]string) // can't mix data types
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	// don't need the following as we are using struct
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. you will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)                                                       // sleeping for 10 seconds
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName) // instead of printing this function will assign it to a var
	fmt.Println("########################################################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("########################################################")
	wg.Done() // removes the thread that we added in the waiting list
}
