### Expense Tracker

## Description

The Expense Tracker is a simple command-line application written in Go that allows users to record and manage their daily expenses directly from the terminal.

This project demonstrates:

Struct usage in Go

Working with slices

User input handling

Basic validation

Loop control and graceful program exit

Modular function design

The program keeps running until the user chooses to exit and handles invalid input without crashing.

⚙️ How It Works

The application runs in a continuous loop and presents a menu with four options:

Add Expenses

Prompts the user for:

Name

Amount

Category

Validates that:

Name is not empty

Amount is a positive number

Category is not empty

Stores the expense in memory using a slice of structs.

View All Expenses

Displays all recorded expenses.

If no expenses exist, it informs the user.

View Total Amount Spent

Calculates the total of all recorded expenses.

Displays the result formatted to two decimal places.

Exit

Stops the program gracefully by breaking the main loop.


``` project structure

Tracker struct → Represents a single expense entry.

expenses slice → Stores all recorded expenses.

displayMenu() → Shows available options.

addExpense() → Collects and validates expense data.

viewAllExpenses() → Displays stored expenses.

viewTotal() → Calculates total spending.

readInput() → Reads and cleans user input.

main() → Controls program flow.

```

## How to run

```option 1
run directly:
 go run main.go
```
```option 2 
build an excutable:
 go build 
then run : 
 ./Expense-tracker