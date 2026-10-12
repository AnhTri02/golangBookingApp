package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	var bookings []string //slice (dynamisk array)

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for remainingTickets > 0 && len(bookings) < 50 {

		firstName, lastName, email, userTickets := getUserInput()
		isValidNames, isValidEmail, isValidTicket := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidNames && isValidEmail && isValidTicket {
			bookTicket(remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName)
			firstNames := getFirstNames(bookings)
			fmt.Printf("The first names of bookings are: %v \n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Conference booked out")
				break
			}
		} else {
			if !isValidNames {
				fmt.Println("Name invalid")
			}
			if !isValidEmail {
				fmt.Println("Email invalid")
			}
			if !isValidTicket {
				fmt.Println("Ticket amount invalid")
			}
		}

	}
}

func greetUsers(conferenceName string, conferenceTickets uint, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets are available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []string) []string {

	firstNames := []string{} //skapar tom slice av string (alla förnamn sparas här)
	//blank identifier unused variables
	for _, x := range bookings { //loopar igenom bookings (alla för och efternamn)
		var nameSeperate = strings.Fields(x)             //split för och efternamn
		firstNames = append(firstNames, nameSeperate[0]) //lägg till endast förnamn
	}

	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	var isValidNames bool = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail bool = strings.Contains(email, "@")
	var isValidTicket bool = userTickets <= remainingTickets

	return isValidNames, isValidEmail, isValidTicket
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName) //passing the memory of the userName veriable

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter the numbers of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string, email string, conferenceName string) {

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("There are %v tickets remaining for %v\n", remainingTickets, conferenceName)

}
