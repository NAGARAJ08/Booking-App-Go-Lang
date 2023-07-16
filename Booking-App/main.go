package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
	"sync"
	"time"
)

//fmt - formatted input output

const ConferenceTickets = 50

var ConfName = "Go Conference" //shorthand notation
var remainingTickets uint = 50 //uint does not allow it to set to negative values
var bookings = make([]userData, 0)

type userData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	// fmt.Printf("ConferenceTickets is %T, remainingTickets is %T, ConfName is %T\n", ConferenceTickets, remainingTickets, ConfName)
	for {

		firstname, lastname, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstname, lastname, email, userTickets, remainingTickets)

		if isValidEmail && isValidName && isValidTicketNumber {

			bookTicket(userTickets, firstname, lastname, email)
			wg.Add(1)
			go sendTickets(userTickets, firstname, lastname, email)
			firstNames := getFirstnames()

			fmt.Printf("The First Names of all our bookings are %v \n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Conference is Booked out. Hope to see you again")
				break
			}

		} else {

			if !isValidName {
				fmt.Println("Invalid Name")
			}

			if !isValidEmail {
				fmt.Println("Invalid Email")
			}

			if !isValidTicketNumber {
				fmt.Println("Invalid Number of Tickets")
			}

			fmt.Printf("Invlid input ")
		}

	}
	wg.Wait()

}

func bookTicket(userTickets uint, firstname string, lastname string, email string) {

	remainingTickets = remainingTickets - userTickets

	//create a map fo a user
	var userData = userData{

		firstName:       firstname,
		lastName:        lastname,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of Bookings is: %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will Receive a Confirmation email at %v\n", firstname, lastname, userTickets, email)
	fmt.Println("Please check your inbox for the Confirmation email")

	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, ConfName)

}

func getUserInput() (string, string, string, uint) {

	var firstname string
	var lastname string
	var email string
	var userTickets uint

	//taking inputs

	fmt.Println("Enter your firstname")
	fmt.Scan(&firstname)

	fmt.Println("Enter your lastname")
	fmt.Scan(&lastname)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("Enter the Number of Tickers required")
	fmt.Scan(&userTickets)

	return firstname, lastname, email, userTickets
}

func getFirstnames() []string {

	firstNames := []string{}

	// _  is used to ignore the var that we dont use
	for _, booking := range bookings {

		var names = strings.Fields(booking.firstName)
		firstNames = append(firstNames, names[0])
	}

	return firstNames

}

func greetUsers() {

	fmt.Println("Welcome to ", ConfName, " Booking application")
	fmt.Println("We have Total of ", ConferenceTickets, "tickets and ", remainingTickets, "are still remaining")
	fmt.Println("Get your Tickets here to attend")

}

func sendTickets(userTickets uint, firstname string, lastname string, email string) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstname, lastname)
	fmt.Println("########################")
	fmt.Printf("Sending Ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("########################")
	wg.Done()
}
