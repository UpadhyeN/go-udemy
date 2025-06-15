package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 6.5
	var investment float64 = 1000
	var years float32
	returns := 5.5
	fmt.Print("Investment Calculator Enter your amount here: ")
	// scanning the function for the input
	fmt.Scan(&investment)
	// This is a pointer to the investment variable
	fmt.Print("Enter years of investment: ")
	fmt.Scan(&years)
	fmt.Print("Enter expected returns in percentage: ")
	fmt.Scan(&returns)
	futureValues := investment * math.Pow(1+returns/100, float64(years))
	fmt.Printf("Future value of the investment: %.2f\n", futureValues)

	futureValuesadjusted := futureValues / math.Pow(1+inflationRate/100, float64(years))
	fmt.Printf("Future value adjusted for inflation: %.2f\n", futureValuesadjusted)
}
