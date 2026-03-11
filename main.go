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
	fmt.Printf("%p\n", &bookings) // prints: 0xc000014078  ← clean address
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
	bookings[0] = firstName + " " + lastName
	fmt.Println("this is whole bookings array : ", bookings)
	fmt.Println("this is first element of bookings array : ", bookings[0])
	fmt.Println("this is length of bookings array : ", len(bookings))
	fmt.Println("this is capacity of bookings array : ", cap(bookings))
	fmt.Printf("this is type of bookings array : %T\n", bookings) // prints: [50]string if we use %T 
	remainingTickets = remainingTickets - userTickets
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conference)
}
