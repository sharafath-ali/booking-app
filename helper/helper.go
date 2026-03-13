package helper

import "fmt"

func PrintWelcome(conference string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("welcome to our " + conference + " booking system\n")
	fmt.Printf(conference)
	fmt.Printf("\n%s\n", conference)
	fmt.Println("welcome to our ", conference, "booking system\n")
	fmt.Printf("we have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("get your tickets here to attend\n")
}