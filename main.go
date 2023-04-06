package main

import (
	"booking-app-nana/helper"
	"fmt"
	"time"
)

const conferenceName = "Go Conference"
const conferenceTickets uint = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInputs(firstName, lastName, email, int(userTickets), int(remainingTickets))

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)
			go sendTicket(firstName, userTickets)
			firstNames := getFirstNames()

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

func greetUsers() {
	fmt.Printf("conferenceName is %T, conferenceTickets is %T and remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v and %v are still available\n", conferenceTickets, remainingTickets)
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
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

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("%v booked %v tickets using %v email address.\n", userData.firstName, userTickets, email)
	fmt.Printf("%v tickets are remained.\n", remainingTickets)
}

func sendTicket(name string, tickets uint) {
	time.Sleep(5 * time.Second)
	ticketsMsg := fmt.Sprintf("%v tickets are being send.\n", tickets)
	fmt.Printf("###############################")
	fmt.Printf("%v thanks for buying the tickets.\n%v", name, ticketsMsg)
	fmt.Printf("###############################")
}
