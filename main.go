package main

import (
	"fmt" 
	"strings"

)


func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	var bookings []string  //slice (dynamisk array)

	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets are available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	
	for{

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

			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName + " " + lastName)


			fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve confirmation email at %v \n", firstName, lastName, userTickets, email)
			fmt.Printf("There are %v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := [] string{}
			//blank identifier unused variables
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames= append(firstNames, names[0])
			}
		
			fmt.Printf("The first names of bookings are: %v\n", firstNames);
	}
	
	

}
