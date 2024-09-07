package main

import (
	"booking-app/helper"
	"fmt"
	// "strconv"
)

var conferenceName = "Go conference"
const conferenceTickets = 50
var remainingTickets uint = 50
var bookings = make([]userData, 0)

type userData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}


func main() {
	// Function call to greet users
	greetUser()

	// Loop runs as long as there are remaining tickets
	for remainingTickets > 0 {
		// User input function
		firstName, lastName, email, userTickets := getUserInput()

		// Validation function
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		// Only proceed if all inputs are valid
		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(&remainingTickets, firstName, lastName, email, userTickets, &bookings)

			// Print first names
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

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
}

func greetUser() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("Get your tickets here to attend\n")
	fmt.Printf("We have a total of %v tickets and %v are available\n", conferenceTickets, remainingTickets)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// Extract first name from the booking map
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

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

func bookTicket(remainingTickets *uint, firstName string, lastName string, email string, userTickets uint, bookings *[]userData) {
	*remainingTickets -= userTickets

// create a struct for a user
var userData= userData{
	firstName: firstName,
	lastName: lastName,
	email: email,
	numberOfTickets: userTickets,
}



	// Create a map for user data
	// userData := map[string]string{
	// 	"firstName":       firstName,
	// 	"lastName":        lastName,
	// 	"email":           email,
	// 	"numberOfTickets": strconv.FormatUint(uint64(userTickets), 10),
	// }

	*bookings = append(*bookings, userData)
	fmt.Printf("list of booking is %v\n", *bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets with us. We will send a confirmation email to %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for the conference.\n", *remainingTickets)
}
