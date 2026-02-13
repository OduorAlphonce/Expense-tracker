package main

import (
	"bufio"   //
	"fmt"     //for printing
	"os"      //access command line arguments
	"strconv" //string conversion
	"strings" //string manipulation
)

// tracker represents a single entry
type Tracker struct {
	Name     string  //name
	Amount   float64 //amount spent
	Category string  //category of expense
}

// container to store all expenses
var expenses []Tracker

// use buffer to read user input from terminal
var input = bufio.NewReader(os.Stdin)

// shows available options to the user
func displayMenu() {
	fmt.Println("1. Add Expenses")
	fmt.Println("2. View all expenses")
	fmt.Println("3. View total amount spent")
	fmt.Println("4. Exit")
}

// collects user input and stores the expense
func addExpense() {
	// read the input name
	name := readInput("Enter your name: ")
	if name == "" {
		fmt.Println("fill the name section")
		return //exit if invalid
	}

	//read input ammount as string
	money := readInput("Enter amount: 0")

	//convert string to float64
	amount, err := strconv.ParseFloat(money, 64)

	//validate the amount
	if err != nil || amount <= 0 {
		fmt.Println("Amount must positive.")
		return
	}

	//read category
	category := readInput("Enter category: ")
	if category == "" {
		fmt.Println("fill in the category")
		return
	}

	expense := Tracker{
		Name:     name,
		Amount:   amount,
		Category: category,
	}

	expenses = append(expenses, expense)

	fmt.Println("added succesfully")
}

// this function displays all recorded expenses
func viewAllExpenses() {
	//checker whether empty expenses exist
	if len(expenses) == 0 {
		fmt.Println("No recorded expenses.")
		return
	}

	//use range loop to loop through the slice and recorde expense
	for _, e := range expenses {
		fmt.Printf("%s, .%.2f, %s", e.Name, e.Amount, e.Category)
	}
}

// viewTotal calculates
func viewTotal() {
	// If no expenses, total is zero
	if len(expenses) == 0 {
		fmt.Println("Total spent: 0.00")
		return
	}

	var total float64

	// Sum all expense amounts
	for _, e := range expenses {
		total += e.Amount
	}

	// Display total formatted to 2 decimal places
	fmt.Printf("Total spent: %.2f\n", total)
}

// readInput prints a prompt and reads user input from the terminal
func readInput(prompt string) string {
	fmt.Print(prompt)

	// Read until newline
	file, _ := input.ReadString('\n')

	// Remove newline and surrounding whitespace
	return strings.TrimSpace(file)
}

func main() {
	// boolen experession controls whether the program keeps running
	running := true

	// Main program loop (runs until user exits)
	for running {
		displayMenu() // Show menu options

		// Read user choice
		choice := readInput("Choose an option: ")

		// Handle menu selection
		switch choice {
		case "1":
			addExpense() // Add a new expense
		case "2":
			viewAllExpenses() // Display all expenses
		case "3":
			viewTotal() // Show total amount spent
		case "4":
			fmt.Println("Exit!")
			running = false // Stop the loop → graceful exit
		default:
			// Handles invalid menu choices
			fmt.Println("Invalid option. Please choose between 1 and 4.")
		}

		// Print empty line for cleaner formatting
		fmt.Println()
	}
}
