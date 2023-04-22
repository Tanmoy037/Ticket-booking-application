package main

import (
    "fmt"
    "strings"
)

func main() {
    conferenceName := "Go Conference"
    var conferenceTickets uint = 50
    var remainingTickets uint = 50
    bookings := []string{}

    fmt.Printf("Welcome %v booking application ", conferenceName )

    fmt.Println("We have total of", conferenceTickets ,"tickets and", remainingTickets, "are still available")
    fmt.Println("Get your tickets here to attend")

    for {
        var firstName string
        var lastName string
        var email string
        var userTicket int
        fmt.Println("Enter your first name: ")
        fmt.Scan(&firstName)
        fmt.Println("Enter your last name: ")
        fmt.Scan(&lastName)
        fmt.Println("Enter your email: ")
        fmt.Scan(&email)
        fmt.Println("Enter number of tickets: ")
        _, err := fmt.Scan(&userTicket)
        if err != nil {
            fmt.Printf("Invalid input for number of tickets: %v ", err)
            continue
        }
        if userTicket < 0 {
            fmt.Println("You cannot book a negative number of tickets")
            continue
        }

        if uint(userTicket) > remainingTickets {
            fmt.Printf("We only have %v tickets remaining. You cannot book %v tickets. \n", remainingTickets, userTicket)
            continue
        }

        remainingTickets -= uint(userTicket)
        bookings = append(bookings, firstName + " " + lastName)

        fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTicket, email)
        fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)

        firstNames := []string{}
        for _, booking := range bookings {
            names := strings.Fields(booking)
            firstNames = append(firstNames, names[0])
        }

        fmt.Printf("These first names of bookings are: %v ", firstNames)

        if remainingTickets == 0 {
            fmt.Println("Our conference is booked out. Come back next year.")
            break
        }
    }
}
