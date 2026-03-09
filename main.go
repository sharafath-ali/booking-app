package main

import "fmt"

func main() {
	var conference string = "Go Conference"
	fmt.Printf("welcome to our " + conference + " booking system\n")
	fmt.Printf(conference)
	fmt.Printf("\n%s\n", conference)
	fmt.Printf("welcome to our %s booking system\n", conference)
	fmt.Printf("please book your tickets for %v\n", conference)
}