package main

import (
	"fmt"
)

const conferenceTickets int = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberofTickets uint
}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)
		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("first names: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our meeting is booked out!!")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("First Name or Last Name is too short")
			}
			if !isValidEmail {
				fmt.Println("Email invalid")
			}
			if !isValidTicketNumber {
				fmt.Println("Ticket Number wrong")
			}

		}

	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("We have a total of %v and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Please follow the instructions to book your tickets.")
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

	fmt.Println("Enter your First name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your Last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your Email address:")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberofTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("List of bookimgs is %v\n", bookings)

	fmt.Printf("Thanks %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("There are %v tickets remaining for %v.\n", remainingTickets, conferenceName)
}
