package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	var conferenceTicktes = 50
	var remainingTicktes = 50

	fmt.Printf("Welcome %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTicktes ,"tickets and", remainingTicktes, "are still available")
	fmt.Println("Get your ticktes here to attend")

	var firstName string
	var lastName string
	var email string
	var userTicket int
	fmt.Println("Enter first your name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTicket)


	

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v .\n", firstName, lastName, userTicket, email)


}