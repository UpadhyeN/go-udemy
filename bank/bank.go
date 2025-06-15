package main

import (
	"fmt"
)

func main() {
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

}
