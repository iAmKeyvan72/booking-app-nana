package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceName is %T, conferenceTickets is %T and remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v and %v are still available\n", conferenceTickets, remainingTickets)

	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address")
	fmt.Scan(&email)

	fmt.Println("How many tickets you need?")
	fmt.Scan(&userTickets)

	fmt.Printf("%v %v booked %v tickets using %v email address.\n", firstName, lastName, userTickets, email)
}
