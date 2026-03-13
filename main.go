package main

import (
	"fmt"
	"strings"
	"bookingapp/helper"
)

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	for {
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		if len(firstName) < 2 || !isValidName(firstName) {
			fmt.Println("Invalid input! Please enter a valid first name.")
			continue // ✅ inside for loop — valid
		}
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		if len(lastName) < 2 || !isValidName(lastName) {
			fmt.Println("Invalid input! Please enter a valid last name.")
			continue // ✅ inside for loop — valid
		}

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		if !strings.Contains(email, "@") {
			fmt.Println("Invalid input! Please enter a valid email.")
			continue // ✅ inside for loop — valid
		}

		break // all valid — exit the loop
	}

	return firstName, lastName, email, userTickets
}

// package level varible
var conference string = "Go Conference"

func main() {
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookingsArray = [50]string{}
	var bookingsSlice = []string{}
	helper.PrintWelcome(conference, conferenceTickets, remainingTickets)
	for remainingTickets > 0 {
		// list of total booked users array
		// pointer in go
		fmt.Printf("%v\n", bookingsArray)
		fmt.Println(bookingsArray)
		fmt.Printf("%p\n", &bookingsArray) // prints: 0xc000014078  ← clean address

		firstName, lastName, email, userTickets := getUserInput()

		fmt.Println(firstName, lastName, email)
		fmt.Println(&firstName, &lastName, &email)
		fmt.Println("Enter number of tickets: ")
		_, err := fmt.Scan(&userTickets)
		if err != nil {
			fmt.Println("Invalid input! Please enter a number.")
			fmt.Scanln() // clear the input buffer
			continue
		}

		if userTickets == 0 {
			fmt.Println("Please enter at least 1 ticket.")
			continue
		}

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
		remainingTickets -= userTickets // update BEFORE printing so all output shows correct count
		firstnames := Printfirstname(bookingsSlice, firstName, lastName, userTickets, email, conference)
		fmt.Println("this is firstnames : ", firstnames)
		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conference)
		checkRemainingTickets(remainingTickets)
	}
}

func checkRemainingTickets(remainingTickets uint) {
	// if else case
	if remainingTickets == 0 {
		fmt.Println("all tickets are booked")
	} else {
		fmt.Printf("%v tickets are still available\n", remainingTickets)
	}
}

func Printfirstname(bookingsSlice []string, firstName string, lastName string, userTickets uint, email string, conference string) []string {
	var firstnames = []string{}
	// for loop to get firstnames
	for _, items := range bookingsSlice {
		var names = strings.Fields(items)
		firstnames = append(firstnames, names[0])
	}
	return firstnames
}
