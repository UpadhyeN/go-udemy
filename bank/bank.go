package main

import (
	"fmt"
	"os"
)

func WritebalanceToFile(balance float64) error {
	balanceData := fmt.Sprint(balance)
	return os.WriteFile("balance.txt", []byte(balanceData), 0644)

}

func main() {

	var balance float64 = 1000.00 // initial balance
	// for will be infinite loop
	for {
		fmt.Println("Welcome to Bank Application")
		fmt.Println("Please select an option from the menu below:")

		fmt.Println("1. check balance")
		fmt.Println("2. deposit")
		fmt.Println("3. withdraw")
		fmt.Println("4. exit")
		var choice int
		fmt.Print("Enter your choice: ") // taking input from the user
		fmt.Scanln(&choice)
		fmt.Println("You selected option:", choice)

		switch choice {
		case 1:
			fmt.Println("Checking balance...", balance)
			WritebalanceToFile(balance)
		case 2:
			fmt.Print("Enter the amount to deposit: ")
			var depositeAmount float64
			fmt.Scanln(&depositeAmount)
			if depositeAmount <= 0 {
				fmt.Println("Deposit amount cannot be negative!")
				continue
			}
			balance = balance + depositeAmount
			fmt.Println("New balance after deposit is:", balance)
			WritebalanceToFile(balance)
		case 3:
			fmt.Print("Enter the amount to withdraw: ")
			var withdrawAmount float64
			fmt.Scanln(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Withdraw amount cannot be negative!")
				continue
			}
			if withdrawAmount > balance {
				fmt.Println("Insufficient balance!")
			} else {
				balance = balance - withdrawAmount
				fmt.Println("New balance after withdrawal is:", balance)
			}
			WritebalanceToFile(balance)
		case 4:
			fmt.Println("Exiting the application. Thank you!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
