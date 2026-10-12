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

	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets are available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {

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

		if userTickets <= remainingTickets {
				remainingTickets = remainingTickets - userTickets
				bookings = append(bookings, firstName+" "+lastName)

				fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve confirmation email at %v \n", firstName, lastName, userTickets, email)
				fmt.Printf("There are %v tickets remaining for %v\n", remainingTickets, conferenceName)

				firstNames := []string{} //skapar tom slice av string (alla förnamn sparas här)
				//blank identifier unused variables
				
				for _, x := range bookings { //loopar igenom bookings (alla för och efternamn)
					var nameSeperate = strings.Fields(x)             //split för och efternamn
					firstNames = append(firstNames, nameSeperate[0]) //lägg till endast förnamn
				}

				fmt.Printf("The first names of bookings are: %v\n", firstNames)

				if remainingTickets == 0 {
					fmt.Println("Conference booked out")
					break
				}
		} else{
				fmt.Printf("We only have %v tickets remaining, can't book more\n", remainingTickets)
				continue
		}
		

	}

}
