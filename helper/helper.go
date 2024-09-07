package helper

import "strings"
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint,remainingTickets uint ) (bool, bool, bool) {
	// Input validation
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	// debugging line
	// fmt.Printf("isValidName: %v, isValidEmail: %v, isValidTicketNumber: %v\n", isValidName, isValidEmail, isValidTicketNumber)

	return isValidName, isValidEmail, isValidTicketNumber
}
