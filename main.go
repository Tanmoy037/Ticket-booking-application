package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	var conferenceTicktes uint = 50
	var remainingTicktes uint = 50

	fmt.Printf("Welcome %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTicktes ,"tickets and", remainingTicktes, "are still available")
	fmt.Println("Get your ticktes here to attend")

	var bookings []string
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


	remainingTicktes = remainingTicktes - uint(userTicket)
	bookings = append(bookings, firstName + " " + lastName)


	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTicket, email)
	fmt.Printf("%v tickets remaining for %v \n", remainingTicktes, conferenceName)
	fmt.Printf("These are all our bookings %v \n", bookings)

}