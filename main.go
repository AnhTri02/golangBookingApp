package main

import "fmt" //format package

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets are available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	//ask user for name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName) //passing the memory of the userName veriable

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter the numbers of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("There are %v tickets remaining for %v\n", remainingTickets, conferenceName)

}
