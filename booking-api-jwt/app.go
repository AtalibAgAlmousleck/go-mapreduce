package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets int = 50

	fmt.Printf("ConferenceName is %T\n", conferenceName)

    
	fmt.Printf("Welcome to %v Booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Got your ticket here to attend the %v", conferenceName)
    fmt.Println("Hello World")

	var userName string
	var userTicket int

	// user input
	fmt.Println("Please enter your username: ")
	fmt.Scan(&userName)
	fmt.Println("Please enter the number of tickets: ")
	fmt.Scan(&userTicket)

	fmt.Printf("User %v booked %v numbers of tickects.\n", userName, userTicket)
}