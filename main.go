package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceName is %T, conferenceTickets is %T and remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v and %v are still available\n", conferenceTickets, remainingTickets)

	bookings := []string{}

	for {

		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address")
		fmt.Scan(&email)

		fmt.Println("How many tickets you need?")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets remaining\n", remainingTickets)
			break
		}

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("%v booked %v tickets using %v email address.\n", bookings[0], userTickets, email)
		fmt.Printf("%v tickets are remained.\n", remainingTickets)

		firstNames := []string{}

		for _, booking := range bookings {
			names := strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("First names of the bookings are: %v\n", firstNames)
		fmt.Printf("the array has %v items\n", len(bookings))

		noTickets := remainingTickets == 0

		if noTickets {
			fmt.Println("Booked out!")
			break
		}
	}

}
