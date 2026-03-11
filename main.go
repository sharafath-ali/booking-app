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
	var bookingsArray = [50]string{}
	var bookingsSlice = []string{}
	for remainingTickets > 0 {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// list of total booked users array
		// pointer in go
		fmt.Printf("%v\n", bookingsArray)
		fmt.Println(bookingsArray)
		fmt.Printf("%p\n", &bookingsArray) // prints: 0xc000014078  ← clean address
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

		if userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets left! Please try again.\n", remainingTickets)
			continue
		}

		bookingsSlice = append(bookingsSlice, firstName+" "+lastName)
		fmt.Println("this is whole bookings slice : ", bookingsSlice)
		fmt.Println("this is first element of bookings slice : ", bookingsSlice[0])
		fmt.Println("this is length of bookings slice : ", len(bookingsSlice))
		fmt.Println("this is capacity of bookings slice : ", cap(bookingsSlice))
		fmt.Printf("this is type of bookings slice : %T\n", bookingsSlice) // prints: []string if we use %T
		// bookingsArray[0] = firstName + " " + lastName
		// fmt.Println("this is whole bookings array : ", bookingsArray)
		// fmt.Println("this is first element of bookings array : ", bookingsArray[0])
		// fmt.Println("this is length of bookings array : ", len(bookingsArray))
		// fmt.Println("this is capacity of bookings array : ", cap(bookingsArray))
		// fmt.Printf("this is type of bookings array : %T\n", bookingsArray) // prints: [50]string if we use %T
		remainingTickets = remainingTickets - userTickets
		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conference)
	}
}
