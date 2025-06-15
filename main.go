package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64
	fmt.Print("Enter the revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter the expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter the tax rate (as a percentage): ")
	fmt.Scan(&taxRate)

	EBT := revenue - expenses
	fmt.Printf("The Earnings Before Tax (EBT) is: %.2f\n", EBT)
	TaxrateDecimal := taxRate / 100
	profit := EBT * (1 - TaxrateDecimal)
	fmt.Printf("The profit after tax is: %.2f\n", profit)
	ratio := profit / revenue
	fmt.Printf("The profit margin is: %.2f%%\n", ratio*100)

}
