package main

import (
	"fmt"
	"golangBookingApp/helper"
	"sync"
	"time"
)

const conferenceTickets uint = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0) //initize a lsit of UserData struct

// som en lightweight klass i java
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}


var wg = sync.WaitGroup{}


func main() {

	greetUsers()



		firstName, lastName, email, userTickets := getUserInput()
		isValidNames, isValidEmail, isValidTicket := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidNames && isValidEmail && isValidTicket {
			bookTicket(userTickets, firstName, lastName, email)


			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v \n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Conference booked out")
				//break
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
		wg.Wait()
	}


func greetUsers() {
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets are available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, s := range bookings {
		firstNames = append(firstNames, s.firstName)
	}

	return firstNames
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

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("There are %v tickets remaining for %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	time.Sleep(5 * time.Second)
	fmt.Println("###################")
	fmt.Printf("Sending ticket:\n %v\n to email adress: %v\n ", ticket, email)
	fmt.Println("###################")

	wg.Done() //remove thread från Add1

}
