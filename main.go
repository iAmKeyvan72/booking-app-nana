package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	greetUsers(conferenceName, int(conferenceTickets), int(remainingTickets))

	bookings := []string{}

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInputs(firstName, lastName, email, int(userTickets), int(remainingTickets))

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets, bookings := bookTicket(remainingTickets, userTickets, bookings, firstName, lastName, email)
			firstNames := getFirstNames(bookings)

			fmt.Printf("First names of the bookings are: %v\n", firstNames)
			fmt.Printf("the array has %v items\n", len(bookings))

			if remainingTickets == 0 {
				fmt.Println("Booked out!")
				break
			}
		} else {
			fmt.Printf("No a valid info provided\n")
		}
	}
}

func greetUsers(conferenceName string, conferenceTickets int, remainingTickets int) {
	fmt.Printf("conferenceName is %T, conferenceTickets is %T and remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v and %v are still available\n", conferenceTickets, remainingTickets)
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}

	for _, booking := range bookings {
		names := strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func validateUserInputs(firstName string, lastName string, email string, userTickets int, remainingTickets int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string, email string) (uint, []string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("%v booked %v tickets using %v email address.\n", bookings[0], userTickets, email)
	fmt.Printf("%v tickets are remained.\n", remainingTickets)

	return remainingTickets, bookings
}
