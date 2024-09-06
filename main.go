package main

import (
	"fmt"
	"strings"
)

func main() {

		var conferenceName = "Go conference"
		// conferenceName :="go conference"   we can also use this syntax or as they call it syntactic sugar and it only works for variables abd not for constants and we can not assign a type explicitly
		const conferenceTickets = 50
		var remainingTickets uint=50
		

	// fmt.Println("hello welcome to", conferenceName, "booking application")
	fmt.Printf("welcome to our %v booking application\n ", conferenceName)
	fmt.Printf("get ur thickets here to attend")
	fmt.Printf(" we have total of %v thickets and  %v are available \n", conferenceTickets, remainingTickets)

	fmt.Printf(" conferenceTicker is %T, remainingTicket %T conferenceName is %T\n",conferenceTickets,remainingTickets, conferenceName )

	var bookings = []string{}

	for {

			// arrays and slices
		// var bookings [50] string

		// slice is a list that is dynamic in size
		


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
		fmt.Printf("how many tickets do u want? ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets-userTickets

		// bookings[0] = firstName + " " + lastName this is the array way of doing things and it is not dynamic in size n also we should keep track of the indices in order to append ..n that is inconvenient

		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("the whole array : %v\n", bookings)
		fmt.Printf("the first value %v\n", bookings[0])
		fmt.Printf("the slice type %T\n", bookings)
		fmt.Printf("the slice length %v\n", len(bookings))

		fmt.Printf("thank you %v %v for booking %v thickets with us.we will send a confirmation email via %v \n ", firstName, lastName, userTickets,email)
		fmt.Printf("%v tickets are remaining for the conference\n", remainingTickets)

		firstNames :=[]string{}

		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

          // note 
			// strings.Fields() splits the string white spaces and returns a slice of strings

		fmt.Printf(" the first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("our conference is booked out. come back next year")
			break
		}
		
	}


	

	

}
