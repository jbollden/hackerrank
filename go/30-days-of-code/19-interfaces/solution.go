// Day 19: Interfaces (https://www.hackerrank.com/challenges/30-interfaces)
// Since Go is not a selectable language this solution has not been tested on HackerRank

package main

import "fmt"

type AdvancedArithmetic interface {
	divisorSum(int) int
}

type Calculator struct{}

func (c Calculator) divisorSum(n int) (sum int) {
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}

// Validating if Calculator has implemented AdvancedArithmetic, if it
// isn't compilation will fail
var _ AdvancedArithmetic = Calculator{}

func main() {
	var n int
	fmt.Scan(&n)
	myCalculator := Calculator{}
	if true {
		sum := myCalculator.divisorSum(n)
		fmt.Printf("I implemented: AdvancedArithmetic\n%v\n", sum)
	} else {
		fmt.Println("Wrong answer") // Will never happen
	}
}
