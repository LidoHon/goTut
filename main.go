package main

import (
	"fmt"
	"strings"
	"booking-app/helper"
)


var conferenceName = "Go conference"
const conferenceTickets = 50
var remainingTickets uint = 50
var bookings = []string{}
func main() {



	// function call to greet users
	greetUser()
	

	// slice is a list that is dynamic in size
	

	// Loop runs as long as there are remaining tickets
	for remainingTickets > 0 {

		// user input func
		firstName, lastName, email, userTickets := getUserInput()

		// validation func
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		// Only proceed if all inputs are valid
		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(&remainingTickets, firstName, lastName, email, userTickets, &bookings)

			// bookings[0] = firstName + " " + lastName
			// This is the array way of doing things, and it is not dynamic in size.
			// We would need to keep track of the indices in order to append, and that is inconvenient.

			// Print slice details
			// fmt.Printf("The whole array: %v\n", bookings)
			// fmt.Printf("The first value: %v\n", bookings[0])
			// fmt.Printf("The slice type: %T\n", bookings)
			// fmt.Printf("The slice length: %v\n", len(bookings))

			// function call to print first names
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			// Extract first names of users from the bookings

			// Check if all tickets are sold out
			if remainingTickets == 0 {
				fmt.Println("Our conference is fully booked. Come back next year!")
				break
			}
		} else {
			// Handle invalid input
			if !isValidName {
				fmt.Println("First name or last name is too short. Please try again!")
			}
			if !isValidEmail {
				fmt.Println("Email address is invalid. Please try again!")
			}

			if !isValidTicketNumber {
				fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets. Please try again!\n", remainingTickets, userTickets)
			}
			continue
		}
	}

	// // switch statement
	// 	city := "London"
	// 	switch city {
	// 		case "New York":
	// 			fmt.Println("Welcome to New York")
	// 		case "Singapore", "London":
	// 			fmt.Println("Welcome to Singapore")
	// 		case "tokyo":
	// 			fmt.Println("Welcome to London")
	// 		default:
	// 			fmt.Println("No city found")
	// 	}
}

func greetUser() {
	// conferenceName :="golang"
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("Get your tickets here to attend\n")
	fmt.Printf("We have a total of %v tickets and %v are available\n", conferenceTickets, remainingTickets)

	// %T prints the type of a variable
	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking) // strings.Fields() splits the string at white spaces and returns a slice of strings
		firstNames = append(firstNames, names[0])

		// Print first names of all bookings
	}
	return firstNames
}


func getUserInput() (string, string, string, uint) {
	// Variables to capture user details
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask a user for their name
	fmt.Printf("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email: ")
	fmt.Scan(&email)

	fmt.Printf("How many tickets do you want? ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(remainingTickets *uint, firstName string, lastName string, email string, userTickets uint, bookings *[]string) {
	// Update remaining tickets
	*remainingTickets = *remainingTickets - userTickets
	*bookings = append(*bookings, firstName+" "+lastName)
	// Confirmation message for the user
	fmt.Printf("Thank you %v %v for booking %v tickets with us. We will send a confirmation email to %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for the conference.\n", *remainingTickets)
}
