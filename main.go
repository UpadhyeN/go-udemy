package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 6.5
	var investment float64 = 1000
	years := 10
	returns := 5.5
	futureValues := investment * math.Pow(1+returns/100, float64(years))
	fmt.Printf("Future value of the investment: %.2f\n", futureValues)

	futureValuesadjusted := futureValues / math.Pow(1+inflationRate/100, float64(years))
	fmt.Printf("Future value adjusted for inflation: %.2f\n", futureValuesadjusted)

}
