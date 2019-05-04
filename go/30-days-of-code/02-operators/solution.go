// Day 2: Operators (https://www.hackerrank.com/challenges/30-operators)

package main

import "fmt"

func round(val float64) int {
	if val < 0 {
		return int(val - 0.5)
	}
	return int(val + 0.5)
}

func main() {

	var mealCost float64
	var tipPercent, taxPercent int

	fmt.Scan(&mealCost)
	fmt.Scan(&tipPercent)
	fmt.Scan(&taxPercent)

	tip := mealCost * float64(tipPercent) / 100
	tax := mealCost * float64(taxPercent) / 100

	totalCost := mealCost + tip + tax
	fmt.Printf("The total meal cost is %v dollars.\n", round(totalCost))
}
