package main

import (
	"fmt"
	"starter/helper"
	"sync"
	"time"
)

var bookings = []UserData{}

var conferenceName = "Glxymesh Conference"

const ticketCount uint = 50

var remainingTickets uint = 50

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	helper.PrintGreetingFromHelper()
	greetUser()
	printRemainingTickets()

	for {

		userData := getUserInput()

		if userData.numberOfTickets > remainingTickets {
			fmt.Printf("Sorry, we only have %v tickets remaining\n", remainingTickets)
			continue
		}

		remainingTickets -= userData.numberOfTickets
		addBooking(conferenceName, userData)

		wg.Add(1)
		go sendTickets(userData)

		firstNames := getFirstNames()
		fmt.Printf("Bookings: %v\n", firstNames)

		printRemainingTickets()

		if remainingTickets == 0 {
			fmt.Println("Sorry, we are sold out")
			break
		}
	}

	wg.Wait()
}

func greetUser() {
	fmt.Printf("Welcome to the %v to all\n", conferenceName)
	fmt.Println("Please get your ticket from the panel window")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() UserData {
	var firstName string
	var lastName string
	var numberOfTickets uint
	var email string

	fmt.Print("Please enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Please enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email address:")
	fmt.Scan(&email)
	fmt.Print("How many tickets would you like to buy?\n")
	fmt.Scan(&numberOfTickets)
	return UserData{firstName, lastName, email, numberOfTickets}
}

func printRemainingTickets() {
	fmt.Printf("We have total of %v tickets remaining out of %v tickets\n", remainingTickets, ticketCount)
}

func addBooking(conferenceName string, userData UserData) {
	bookings = append(bookings, userData)
	fmt.Printf("Welcome to the %v to %v, You have %v tickets.\n", conferenceName, userData.firstName, userData.numberOfTickets)
	fmt.Println("You will receive an email with your ticket details")
	fmt.Printf("List of bookings: %v\n", bookings)
}

func sendTickets(userData UserData) {
	time.Sleep(50 * time.Second)
	var message = fmt.Sprintf("Sending tickets to %v %v at %v for %v tickets", userData.firstName, userData.lastName, userData.email, userData.numberOfTickets)
	fmt.Println("\n--*--*--*--*--*--*--*--*--*--*--*--*--*--*--*--")
	fmt.Println("Sending tickets...")
	fmt.Println(message)
	fmt.Println("--*--*--*--*--*--*--*--*--*--*--*--*--*--*--*--")
	fmt.Println()
	wg.Done()
}
