package main

import "fmt"

func main() {
	var conference string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	fmt.Printf("welcome to our " + conference + " booking system\n")
	fmt.Printf(conference)
	fmt.Printf("\n%s\n", conference)
	fmt.Println("welcome to our ", conference, "booking system\n")
	fmt.Printf("we have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("get your tickets here to attend\n")
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// list of total booked users array
	var bookings = [50]string{}
	// pointer in go 
	fmt.Printf("%v\n", bookings)
	fmt.Println(bookings)
	fmt.Println(&bookings)
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println(firstName, lastName, email)
	fmt.Println(&firstName, &lastName, &email)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)
	remainingTickets = remainingTickets - userTickets
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conference)
}
