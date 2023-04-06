package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceName is %T, conferenceTickets is %T and remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v and %v are still available\n", conferenceTickets, remainingTickets)

	bookings := []string{}

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

	remainingTickets = conferenceTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("%v booked %v tickets using %v email address.\n", bookings[0], userTickets, email)
	fmt.Printf("%v tickets are remained.\n", remainingTickets)
	fmt.Printf("%v\n", len(bookings))
}
